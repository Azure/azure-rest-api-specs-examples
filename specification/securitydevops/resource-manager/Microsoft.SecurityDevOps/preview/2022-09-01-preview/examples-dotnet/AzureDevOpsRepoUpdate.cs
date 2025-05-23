using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityDevOps;

// Generated from example definition: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoUpdate.json
// this example is just showing the usage of "AzureDevOpsRepo_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AzureDevOpsRepoResource created on azure
// for more information of creating AzureDevOpsRepoResource, please refer to the document of AzureDevOpsRepoResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "westusrg";
string azureDevOpsConnectorName = "testconnector";
string azureDevOpsOrgName = "myOrg";
string azureDevOpsProjectName = "myProject";
string azureDevOpsRepoName = "myRepo";
ResourceIdentifier azureDevOpsRepoResourceId = AzureDevOpsRepoResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureDevOpsConnectorName, azureDevOpsOrgName, azureDevOpsProjectName, azureDevOpsRepoName);
AzureDevOpsRepoResource azureDevOpsRepo = client.GetAzureDevOpsRepoResource(azureDevOpsRepoResourceId);

// invoke the operation
AzureDevOpsRepoData data = new AzureDevOpsRepoData();
ArmOperation<AzureDevOpsRepoResource> lro = await azureDevOpsRepo.UpdateAsync(WaitUntil.Completed, data);
AzureDevOpsRepoResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AzureDevOpsRepoData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
