using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Confluent;
using Azure.ResourceManager.Confluent.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_CreateClusterAPIKey.json
// this example is just showing the usage of "Organization_CreateAPIKey" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConfluentOrganizationResource created on azure
// for more information of creating ConfluentOrganizationResource, please refer to the document of ConfluentOrganizationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string organizationName = "myOrganization";
ResourceIdentifier confluentOrganizationResourceId = ConfluentOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName);
ConfluentOrganizationResource confluentOrganization = client.GetConfluentOrganizationResource(confluentOrganizationResourceId);

// invoke the operation
string environmentId = "env-12132";
string clusterId = "clusterId-123";
ConfluentApiKeyCreateContent content = new ConfluentApiKeyCreateContent()
{
    Name = "CI kafka access key",
    Description = "This API key provides kafka access to cluster x",
};
ConfluentApiKeyRecord result = await confluentOrganization.CreateApiKeyAsync(environmentId, clusterId, content);

Console.WriteLine($"Succeeded: {result}");
