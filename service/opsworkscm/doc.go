// Code generated by smithy-go-codegen DO NOT EDIT.

// Package opsworkscm provides the API client, operations, and parameter types for
// AWS OpsWorks CM.
//
// # OpsWorks CM
//
// The OpsWorks services have reached end of life and have been disabled for both
// new and existing customers. We strongly recommend customers migrate their
// workloads to other solutions as soon as possible. If you have questions about
// migration, reach out to the Amazon Web ServicesSupport Team on [Amazon Web Services re:Post]or through [Amazon Web Services Premium Support].
//
// OpsWorks CM is a service that runs and manages configuration management
// servers. You can use OpsWorks CM to create and manage OpsWorks for Chef Automate
// and OpsWorks for Puppet Enterprise servers, and add or remove nodes for the
// servers to manage.
//
// Glossary of terms
//
//   - Server: A configuration management server that can be highly-available. The
//     configuration management server runs on an Amazon Elastic Compute Cloud (EC2)
//     instance, and may use various other Amazon Web Services services, such as Amazon
//     Relational Database Service (RDS) and Elastic Load Balancing. A server is a
//     generic abstraction over the configuration manager that you want to use, much
//     like Amazon RDS. In OpsWorks CM, you do not start or stop servers. After you
//     create servers, they continue to run until they are deleted.
//
//   - Engine: The engine is the specific configuration manager that you want to
//     use. Valid values in this release include ChefAutomate and Puppet .
//
//   - Backup: This is an application-level backup of the data that the
//     configuration manager stores. OpsWorks CM creates an S3 bucket for backups when
//     you launch the first server. A backup maintains a snapshot of a server's
//     configuration-related attributes at the time the backup starts.
//
//   - Events: Events are always related to a server. Events are written during
//     server creation, when health checks run, when backups are created, when system
//     maintenance is performed, etc. When you delete a server, the server's events are
//     also deleted.
//
//   - Account attributes: Every account has attributes that are assigned in the
//     OpsWorks CM database. These attributes store information about configuration
//     limits (servers, backups, etc.) and your customer account.
//
// # Endpoints
//
// OpsWorks CM supports the following endpoints, all HTTPS. You must connect to
// one of the following endpoints. Your servers can only be accessed or managed
// within the endpoint in which they are created.
//
//   - opsworks-cm.us-east-1.amazonaws.com
//
//   - opsworks-cm.us-east-2.amazonaws.com
//
//   - opsworks-cm.us-west-1.amazonaws.com
//
//   - opsworks-cm.us-west-2.amazonaws.com
//
//   - opsworks-cm.ap-northeast-1.amazonaws.com
//
//   - opsworks-cm.ap-southeast-1.amazonaws.com
//
//   - opsworks-cm.ap-southeast-2.amazonaws.com
//
//   - opsworks-cm.eu-central-1.amazonaws.com
//
//   - opsworks-cm.eu-west-1.amazonaws.com
//
// For more information, see [OpsWorks endpoints and quotas] in the Amazon Web Services General Reference.
//
// # Throttling limits
//
// All API operations allow for five requests per second with a burst of 10
// requests per second.
//
// [OpsWorks endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/opsworks-service.html
// [Amazon Web Services re:Post]: https://repost.aws/
// [Amazon Web Services Premium Support]: https://aws.amazon.com/support
package opsworkscm
