using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/GetDevOpsOperationResultsSucceeded_example.json
// this example is just showing the usage of "DevOpsOperationResults_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevOpsConfigurationResource created on azure
// for more information of creating DevOpsConfigurationResource, please refer to the document of DevOpsConfigurationResource
string subscriptionId = "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
string resourceGroupName = "myRg";
string securityConnectorName = "mySecurityConnectorName";
ResourceIdentifier devOpsConfigurationResourceId = DevOpsConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, securityConnectorName);
DevOpsConfigurationResource devOpsConfiguration = client.GetDevOpsConfigurationResource(devOpsConfigurationResourceId);

// invoke the operation
string operationResultId = "4e826cf1-5c36-4808-a7d2-fb4f5170978b";
OperationStatusResult result = await devOpsConfiguration.GetDevOpsOperationResultAsync(operationResultId);

Console.WriteLine($"Succeeded: {result}");
