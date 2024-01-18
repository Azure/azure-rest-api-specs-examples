using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPolicy_example.json
// this example is just showing the usage of "JitNetworkAccessPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this JitNetworkAccessPolicyResource created on azure
// for more information of creating JitNetworkAccessPolicyResource, please refer to the document of JitNetworkAccessPolicyResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "myRg1";
AzureLocation ascLocation = new AzureLocation("westeurope");
string jitNetworkAccessPolicyName = "default";
ResourceIdentifier jitNetworkAccessPolicyResourceId = JitNetworkAccessPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ascLocation, jitNetworkAccessPolicyName);
JitNetworkAccessPolicyResource jitNetworkAccessPolicy = client.GetJitNetworkAccessPolicyResource(jitNetworkAccessPolicyResourceId);

// invoke the operation
JitNetworkAccessPolicyResource result = await jitNetworkAccessPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
JitNetworkAccessPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
