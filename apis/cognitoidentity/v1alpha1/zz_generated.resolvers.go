/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-aws/apis/cognitoidentityprovider/v1alpha1"
	v1beta1 "github.com/crossplane-contrib/provider-aws/apis/iam/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IdentityPool.
func (mg *IdentityPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomIdentityPoolParameters.OpenIDConnectProviderARNs),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CustomIdentityPoolParameters.OpenIDConnectProviderARNRefs,
		Selector:      mg.Spec.ForProvider.CustomIdentityPoolParameters.OpenIDConnectProviderARNSelector,
		To: reference.To{
			List:    &v1beta1.OpenIDConnectProviderList{},
			Managed: &v1beta1.OpenIDConnectProvider{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomIdentityPoolParameters.OpenIDConnectProviderARNs")
	}
	mg.Spec.ForProvider.CustomIdentityPoolParameters.OpenIDConnectProviderARNs = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CustomIdentityPoolParameters.OpenIDConnectProviderARNRefs = mrsp.ResolvedReferences

	for i4 := 0; i4 < len(mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders); i4++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ClientID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ClientIDRef,
			Selector:     mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ClientIDSelector,
			To: reference.To{
				List:    &v1alpha1.UserPoolClientList{},
				Managed: &v1alpha1.UserPoolClient{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ClientID")
		}
		mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ClientIDRef = rsp.ResolvedReference

	}
	for i4 := 0; i4 < len(mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders); i4++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ProviderName),
			Extract:      v1alpha1.UserPoolName(),
			Reference:    mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ProviderNameRef,
			Selector:     mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ProviderNameSelector,
			To: reference.To{
				List:    &v1alpha1.UserPoolList{},
				Managed: &v1alpha1.UserPool{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ProviderName")
		}
		mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ProviderName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomIdentityPoolParameters.CognitoIdentityProviders[i4].ProviderNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Status.AtProvider.CognitoIdentityProviders); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Status.AtProvider.CognitoIdentityProviders[i3].ClientID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Status.AtProvider.CognitoIdentityProviders[i3].ClientIDRef,
			Selector:     mg.Status.AtProvider.CognitoIdentityProviders[i3].ClientIDSelector,
			To: reference.To{
				List:    &v1alpha1.UserPoolClientList{},
				Managed: &v1alpha1.UserPoolClient{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Status.AtProvider.CognitoIdentityProviders[i3].ClientID")
		}
		mg.Status.AtProvider.CognitoIdentityProviders[i3].ClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Status.AtProvider.CognitoIdentityProviders[i3].ClientIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Status.AtProvider.CognitoIdentityProviders); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Status.AtProvider.CognitoIdentityProviders[i3].ProviderName),
			Extract:      v1alpha1.UserPoolName(),
			Reference:    mg.Status.AtProvider.CognitoIdentityProviders[i3].ProviderNameRef,
			Selector:     mg.Status.AtProvider.CognitoIdentityProviders[i3].ProviderNameSelector,
			To: reference.To{
				List:    &v1alpha1.UserPoolList{},
				Managed: &v1alpha1.UserPool{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Status.AtProvider.CognitoIdentityProviders[i3].ProviderName")
		}
		mg.Status.AtProvider.CognitoIdentityProviders[i3].ProviderName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Status.AtProvider.CognitoIdentityProviders[i3].ProviderNameRef = rsp.ResolvedReference

	}

	return nil
}
