using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/getPolicyAssignmentWithResourceSelectors.json
// this example is just showing the usage of "PolicyAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this PolicyAssignmentResource
string scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
PolicyAssignmentCollection collection = client.GetGenericResource(scopeId).GetPolicyAssignments();

// invoke the operation
string policyAssignmentName = "CostManagement";
bool result = await collection.ExistsAsync(policyAssignmentName);

Console.WriteLine($"Succeeded: {result}");
