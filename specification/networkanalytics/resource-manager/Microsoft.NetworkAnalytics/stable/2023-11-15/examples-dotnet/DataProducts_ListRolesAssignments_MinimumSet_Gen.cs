using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NetworkAnalytics.Models;
using Azure.ResourceManager.NetworkAnalytics;

// Generated from example definition: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_ListRolesAssignments_MinimumSet_Gen.json
// this example is just showing the usage of "DataProducts_ListRolesAssignments" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataProductResource created on azure
// for more information of creating DataProductResource, please refer to the document of DataProductResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "aoiresourceGroupName";
string dataProductName = "dataproduct01";
ResourceIdentifier dataProductResourceId = DataProductResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataProductName);
DataProductResource dataProduct = client.GetDataProductResource(dataProductResourceId);

// invoke the operation
BinaryData body = BinaryData.FromObjectAsJson(new object());
RoleAssignmentListResult result = await dataProduct.GetRolesAssignmentsAsync(body);

Console.WriteLine($"Succeeded: {result}");
