using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MigrationDiscoverySap.Models;
using Azure.ResourceManager.MigrationDiscoverySap;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_Delete.json
// this example is just showing the usage of "SapInstances_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapInstanceResource created on azure
// for more information of creating SapInstanceResource, please refer to the document of SapInstanceResource
string subscriptionId = "6d875e77-e412-4d7d-9af4-8895278b4443";
string resourceGroupName = "test-rg";
string sapDiscoverySiteName = "SampleSite";
string sapInstanceName = "MPP_MPP";
ResourceIdentifier sapInstanceResourceId = SapInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sapDiscoverySiteName, sapInstanceName);
SapInstanceResource sapInstance = client.GetSapInstanceResource(sapInstanceResourceId);

// invoke the operation
await sapInstance.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
