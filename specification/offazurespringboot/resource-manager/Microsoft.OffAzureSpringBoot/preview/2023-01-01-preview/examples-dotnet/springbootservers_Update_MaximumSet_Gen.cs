using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SpringAppDiscovery;
using Azure.ResourceManager.SpringAppDiscovery.Models;

// Generated from example definition: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootservers_Update_MaximumSet_Gen.json
// this example is just showing the usage of "springbootservers_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SpringBootServerResource created on azure
// for more information of creating SpringBootServerResource, please refer to the document of SpringBootServerResource
string subscriptionId = "etmdxomjncqvygm";
string resourceGroupName = "rgspringbootservers";
string siteName = "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj";
string springbootserversName = "zkarbqnwnxeozvjrkpdqmgnwedwgtwcmmyqwaijkn";
ResourceIdentifier springBootServerResourceId = SpringBootServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName, springbootserversName);
SpringBootServerResource springBootServer = client.GetSpringBootServerResource(springBootServerResourceId);

// invoke the operation
SpringBootServerPatch patch = new SpringBootServerPatch();
ArmOperation<SpringBootServerResource> lro = await springBootServer.UpdateAsync(WaitUntil.Completed, patch);
SpringBootServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SpringBootServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
