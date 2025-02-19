// Code generated by smithy-go-codegen DO NOT EDIT.

// Package cognitoidentityprovider provides the API client, operations, and
// parameter types for Amazon Cognito Identity Provider.
//
// With the Amazon Cognito user pools API, you can configure user pools and
// authenticate users. To authenticate users from third-party identity providers
// (IdPs) in this API, you can [link IdP users to native user profiles]. Learn more about the authentication and
// authorization of federated users at [Adding user pool sign-in through a third party]and in the [User pool federation endpoints and hosted UI reference].
//
// This API reference provides detailed information about API operations and
// object types in Amazon Cognito.
//
// Along with resource management operations, the Amazon Cognito user pools API
// includes classes of operations and authorization models for client-side and
// server-side authentication of users. You can interact with operations in the
// Amazon Cognito user pools API as any of the following subjects.
//
//   - An administrator who wants to configure user pools, app clients, users,
//     groups, or other user pool functions.
//
//   - A server-side app, like a web application, that wants to use its Amazon Web
//     Services privileges to manage, authenticate, or authorize a user.
//
//   - A client-side app, like a mobile app, that wants to make unauthenticated
//     requests to manage, authenticate, or authorize a user.
//
// For more information, see [Using the Amazon Cognito user pools API and user pool endpoints] in the Amazon Cognito Developer Guide.
//
// With your Amazon Web Services SDK, you can build the logic to support
// operational flows in every use case for this API. You can also make direct REST
// API requests to [Amazon Cognito user pools service endpoints]. The following links can get you started with the
// CognitoIdentityProvider client in other supported Amazon Web Services SDKs.
//
// [Amazon Web Services Command Line Interface]
//
// [Amazon Web Services SDK for .NET]
//
// [Amazon Web Services SDK for C++]
//
// [Amazon Web Services SDK for Go]
//
// [Amazon Web Services SDK for Java V2]
//
// [Amazon Web Services SDK for JavaScript]
//
// [Amazon Web Services SDK for PHP V3]
//
// [Amazon Web Services SDK for Python]
//
// [Amazon Web Services SDK for Ruby V3]
//
// [Amazon Web Services SDK for Kotlin]
//
// To get started with an Amazon Web Services SDK, see [Tools to Build on Amazon Web Services]. For example actions and
// scenarios, see [Code examples for Amazon Cognito Identity Provider using Amazon Web Services SDKs].
//
// [Adding user pool sign-in through a third party]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation.html
// [Amazon Web Services SDK for Java V2]: https://sdk.amazonaws.com/java/api/latest/software/amazon/awssdk/services/cognitoidentityprovider/CognitoIdentityProviderClient.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Amazon Cognito user pools service endpoints]: https://docs.aws.amazon.com/general/latest/gr/cognito_identity.html#cognito_identity_your_user_pools_region
// [link IdP users to native user profiles]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html
// [Amazon Web Services SDK for Python]: https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/cognito-idp.html
// [Amazon Web Services SDK for Ruby V3]: https://docs.aws.amazon.com/sdk-for-ruby/v3/api/Aws/CognitoIdentityProvider/Client.html
// [Amazon Web Services SDK for C++]: https://sdk.amazonaws.com/cpp/api/LATEST/aws-cpp-sdk-cognito-idp/html/class_aws_1_1_cognito_identity_provider_1_1_cognito_identity_provider_client.html
// [Tools to Build on Amazon Web Services]: http://aws.amazon.com/developer/tools/
// [User pool federation endpoints and hosted UI reference]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-userpools-server-contract-reference.html
// [Amazon Web Services SDK for PHP V3]: https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cognito-idp-2016-04-18.html
// [Amazon Web Services SDK for .NET]: https://docs.aws.amazon.com/sdkfornet/v3/apidocs/items/CognitoIdentityProvider/TCognitoIdentityProviderClient.html
// [Code examples for Amazon Cognito Identity Provider using Amazon Web Services SDKs]: https://docs.aws.amazon.com/cognito/latest/developerguide/service_code_examples_cognito-identity-provider.html
// [Amazon Web Services Command Line Interface]: https://docs.aws.amazon.com/cli/latest/reference/cognito-idp/index.html#cli-aws-cognito-idp
// [Amazon Web Services SDK for Go]: https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider
// [Amazon Web Services SDK for JavaScript]: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/CognitoIdentityServiceProvider.html
// [Amazon Web Services SDK for Kotlin]: https://sdk.amazonaws.com/kotlin/api/latest/cognitoidentityprovider/aws.sdk.kotlin.services.cognitoidentityprovider/-cognito-identity-provider-client/index.html
package cognitoidentityprovider
