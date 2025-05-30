using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceOSFamily_Get.json
// this example is just showing the usage of "CloudServiceOperatingSystems_GetOSFamily" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudServiceOSFamilyResource created on azure
// for more information of creating CloudServiceOSFamilyResource, please refer to the document of CloudServiceOSFamilyResource
string subscriptionId = "{subscription-id}";
AzureLocation location = new AzureLocation("westus2");
string osFamilyName = "3";
ResourceIdentifier cloudServiceOSFamilyResourceId = CloudServiceOSFamilyResource.CreateResourceIdentifier(subscriptionId, location, osFamilyName);
CloudServiceOSFamilyResource cloudServiceOSFamily = client.GetCloudServiceOSFamilyResource(cloudServiceOSFamilyResourceId);

// invoke the operation
CloudServiceOSFamilyResource result = await cloudServiceOSFamily.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudServiceOSFamilyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
