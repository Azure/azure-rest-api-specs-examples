using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/SecurityPolicies_Get.json
// this example is just showing the usage of "FrontDoorSecurityPolicies_Get" operation, for the dependent resources, they will have to be created separately.

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
FrontDoorSecurityPolicyResource result = await frontDoorSecurityPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorSecurityPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
