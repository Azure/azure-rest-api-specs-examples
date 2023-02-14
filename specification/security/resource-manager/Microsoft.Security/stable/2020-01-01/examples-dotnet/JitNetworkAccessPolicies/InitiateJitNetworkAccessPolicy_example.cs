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

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/InitiateJitNetworkAccessPolicy_example.json
// this example is just showing the usage of "JitNetworkAccessPolicies_Initiate" operation, for the dependent resources, they will have to be created separately.

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
JitNetworkAccessPolicyInitiateContent content = new JitNetworkAccessPolicyInitiateContent(new JitNetworkAccessPolicyInitiateVirtualMachine[]
{
new JitNetworkAccessPolicyInitiateVirtualMachine(new ResourceIdentifier("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),new JitNetworkAccessPolicyInitiatePort[]
{
new JitNetworkAccessPolicyInitiatePort(3389,DateTimeOffset.Parse("placeholder"))
{
AllowedSourceAddressPrefix = "192.127.0.2",
}
})
})
{
    Justification = "testing a new version of the product",
};
JitNetworkAccessRequestInfo result = await jitNetworkAccessPolicy.InitiateAsync(content);

Console.WriteLine($"Succeeded: {result}");
