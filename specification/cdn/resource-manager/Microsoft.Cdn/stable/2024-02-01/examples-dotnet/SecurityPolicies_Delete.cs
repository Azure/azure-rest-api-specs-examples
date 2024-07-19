using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/SecurityPolicies_Delete.json
// this example is just showing the usage of "FrontDoorSecurityPolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorSecurityPolicyResource created on azure
// for more information of creating FrontDoorSecurityPolicyResource, please refer to the document of FrontDoorSecurityPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string securityPolicyName = "securityPolicy1";
ResourceIdentifier frontDoorSecurityPolicyResourceId = FrontDoorSecurityPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, securityPolicyName);
FrontDoorSecurityPolicyResource frontDoorSecurityPolicy = client.GetFrontDoorSecurityPolicyResource(frontDoorSecurityPolicyResourceId);

// invoke the operation
await frontDoorSecurityPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
