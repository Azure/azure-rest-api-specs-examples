using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SpringAppDiscovery;
using Azure.ResourceManager.SpringAppDiscovery.Models;

// Generated from example definition: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_Update_MinimumSet_Gen.json
// this example is just showing the usage of "springbootapps_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SpringBootAppResource created on azure
// for more information of creating SpringBootAppResource, please refer to the document of SpringBootAppResource
string subscriptionId = "jnetwlorzmxpxmcucorv";
string resourceGroupName = "rgspringbootapps";
string siteName = "pdfosfhtemfsaglvwjdyqlyeipucrd";
string springbootappsName = "ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf";
ResourceIdentifier springBootAppResourceId = SpringBootAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName, springbootappsName);
SpringBootAppResource springBootApp = client.GetSpringBootAppResource(springBootAppResourceId);

// invoke the operation
SpringBootAppPatch patch = new SpringBootAppPatch();
ArmOperation<SpringBootAppResource> lro = await springBootApp.UpdateAsync(WaitUntil.Completed, patch);
SpringBootAppResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SpringBootAppData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
