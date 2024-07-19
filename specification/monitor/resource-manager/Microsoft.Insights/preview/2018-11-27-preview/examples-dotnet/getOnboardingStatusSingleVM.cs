using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusSingleVM.json
// this example is just showing the usage of "VMInsights_GetOnboardingStatus" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VmInsightsOnboardingStatusResource created on azure
// for more information of creating VmInsightsOnboardingStatusResource, please refer to the document of VmInsightsOnboardingStatusResource
string resourceUri = "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm";
ResourceIdentifier vmInsightsOnboardingStatusResourceId = VmInsightsOnboardingStatusResource.CreateResourceIdentifier(resourceUri);
VmInsightsOnboardingStatusResource vmInsightsOnboardingStatus = client.GetVmInsightsOnboardingStatusResource(vmInsightsOnboardingStatusResourceId);

// invoke the operation
VmInsightsOnboardingStatusResource result = await vmInsightsOnboardingStatus.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VmInsightsOnboardingStatusData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
