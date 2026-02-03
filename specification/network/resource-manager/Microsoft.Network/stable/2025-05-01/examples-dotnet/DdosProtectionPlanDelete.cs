using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DdosProtectionPlanDelete.json
// this example is just showing the usage of "DdosProtectionPlans_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DdosProtectionPlanResource created on azure
// for more information of creating DdosProtectionPlanResource, please refer to the document of DdosProtectionPlanResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string ddosProtectionPlanName = "test-plan";
ResourceIdentifier ddosProtectionPlanResourceId = DdosProtectionPlanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ddosProtectionPlanName);
DdosProtectionPlanResource ddosProtectionPlan = client.GetDdosProtectionPlanResource(ddosProtectionPlanResourceId);

// invoke the operation
await ddosProtectionPlan.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
