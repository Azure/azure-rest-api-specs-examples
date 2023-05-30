using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Consumption;
using Azure.ResourceManager.Consumption.Models;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostByManagementGroup.json
// this example is just showing the usage of "AggregatedCost_GetByManagementGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string managementGroupId = "managementGroupForTest";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(managementGroupId);
ManagementGroupResource managementGroupResource = client.GetManagementGroupResource(managementGroupResourceId);

// invoke the operation
ConsumptionAggregatedCostResult result = await managementGroupResource.GetAggregatedCostAsync();

Console.WriteLine($"Succeeded: {result}");
