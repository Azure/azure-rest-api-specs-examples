using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2022-06-01/examples/privatelink/WorkspacePrivateLinkResourceGet.json
// this example is just showing the usage of "WorkspacePrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisWorkspaceResource created on azure
// for more information of creating HealthcareApisWorkspaceResource, please refer to the document of HealthcareApisWorkspaceResource
string subscriptionId = "subid";
string resourceGroupName = "testRG";
string workspaceName = "workspace1";
ResourceIdentifier healthcareApisWorkspaceResourceId = HealthcareApisWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
HealthcareApisWorkspaceResource healthcareApisWorkspace = client.GetHealthcareApisWorkspaceResource(healthcareApisWorkspaceResourceId);

// get the collection of this HealthcareApisWorkspacePrivateLinkResource
HealthcareApisWorkspacePrivateLinkResourceCollection collection = healthcareApisWorkspace.GetHealthcareApisWorkspacePrivateLinkResources();

// invoke the operation
string groupName = "healthcareworkspace";
bool result = await collection.ExistsAsync(groupName);

Console.WriteLine($"Succeeded: {result}");
