/*
	GoToSocial
	Copyright (C) GoToSocial Authors admin@gotosocial.org
	SPDX-License-Identifier: AGPL-3.0-or-later

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import { replaceCacheOnMutation, removeFromCacheOnMutation } from "../query-modifiers";
import { gtsApi } from "../gts-api";
import { listToKeyedObject } from "../transforms";
import { AdminAccount, HandleSignupParams, SearchAccountParams } from "../../types/account";

const extended = gtsApi.injectEndpoints({
	endpoints: (build) => ({
		updateInstance: build.mutation({
			query: (formData) => ({
				method: "PATCH",
				url: `/api/v1/instance`,
				asForm: true,
				body: formData,
				discardEmpty: true
			}),
			...replaceCacheOnMutation("instanceV1"),
		}),

		mediaCleanup: build.mutation({
			query: (days) => ({
				method: "POST",
				url: `/api/v1/admin/media_cleanup`,
				params: {
					remote_cache_days: days
				}
			})
		}),

		instanceKeysExpire: build.mutation({
			query: (domain) => ({
				method: "POST",
				url: `/api/v1/admin/domain_keys_expire`,
				params: {
					domain: domain
				}
			})
		}),

		getAccount: build.query<AdminAccount, string>({
			query: (id) => ({
				url: `/api/v1/admin/accounts/${id}`
			}),
			providesTags: (_result, _error, id) => [
				{ type: 'Account', id }
			],
		}),

		searchAccounts: build.query<AdminAccount[], SearchAccountParams>({
			query: (form) => {
				const params = new(URLSearchParams);
				Object.entries(form).forEach(([k, v]) => {
					if (v !== undefined) {
						params.append(k, v);
					}
				});

				let query = "";
				if (params.size !== 0) {
					query = `?${params.toString()}`;
				}

				return {
					url: `/api/v2/admin/accounts${query}`
				};
			},
			providesTags: (res) =>
				res
					? [
						...res.map(({ id }) => ({ type: 'Account' as const, id })),
						{ type: 'Account', id: 'LIST' },
					  ]
					: [{ type: 'Account', id: 'LIST' }],
		}),

		actionAccount: build.mutation<string, { id: string, action: string, reason: string }>({
			query: ({ id, action, reason }) => ({
				method: "POST",
				url: `/api/v1/admin/accounts/${id}/action`,
				asForm: true,
				body: {
					type: action,
					text: reason
				}
			}),
			invalidatesTags: (_result, _error, { id }) => [
				{ type: 'Account', id },
			],
		}),

		handleSignup: build.mutation<AdminAccount, HandleSignupParams>({
			query: ({id, approve_or_reject, ...formData}) => {
				return {
					method: "POST",
					url: `/api/v1/admin/accounts/${id}/${approve_or_reject}`,
					asForm: true,
					body: approve_or_reject === "reject" ?? formData,
				};
			},
			invalidatesTags: (_result, _error, { id }) => [
				{ type: 'Account', id },
			],
		}),

		instanceRules: build.query({
			query: () => ({
				url: `/api/v1/admin/instance/rules`
			}),
			transformResponse: listToKeyedObject<any>("id")
		}),

		addInstanceRule: build.mutation({
			query: (formData) => ({
				method: "POST",
				url: `/api/v1/admin/instance/rules`,
				asForm: true,
				body: formData,
				discardEmpty: true
			}),
			transformResponse: (data) => {
				return {
					[data.id]: data
				};
			},
			...replaceCacheOnMutation("instanceRules"),
		}),

		updateInstanceRule: build.mutation({
			query: ({ id, ...edit }) => ({
				method: "PATCH",
				url: `/api/v1/admin/instance/rules/${id}`,
				asForm: true,
				body: edit,
				discardEmpty: true
			}),
			transformResponse: (data) => {
				return {
					[data.id]: data
				};
			},
			...replaceCacheOnMutation("instanceRules"),
		}),

		deleteInstanceRule: build.mutation({
			query: (id) => ({
				method: "DELETE",
				url: `/api/v1/admin/instance/rules/${id}`
			}),
			...removeFromCacheOnMutation("instanceRules", {
				key: (_draft, rule) => rule.id,
			})
		})
	})
});

export const {
	useUpdateInstanceMutation,
	useMediaCleanupMutation,
	useInstanceKeysExpireMutation,
	useGetAccountQuery,
	useActionAccountMutation,
	useSearchAccountsQuery,
	useLazySearchAccountsQuery,
	useHandleSignupMutation,
	useInstanceRulesQuery,
	useAddInstanceRuleMutation,
	useUpdateInstanceRuleMutation,
	useDeleteInstanceRuleMutation,
} = extended;
