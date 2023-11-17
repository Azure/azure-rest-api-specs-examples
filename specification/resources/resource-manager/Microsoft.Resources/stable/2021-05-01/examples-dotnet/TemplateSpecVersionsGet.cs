using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsGet.json
// this example is just showing the usage of "TemplateSpecVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TemplateSpecResource created on azure
// for more information of creating TemplateSpecResource, please refer to the document of TemplateSpecResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "templateSpecRG";
string templateSpecName = "simpleTemplateSpec";
ResourceIdentifier templateSpecResourceId = TemplateSpecResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, templateSpecName);
TemplateSpecResource templateSpec = client.GetTemplateSpecResource(templateSpecResourceId);

// get the collection of this TemplateSpecVersionResource
TemplateSpecVersionCollection collection = templateSpec.GetTemplateSpecVersions();

// invoke the operation
string templateSpecVersion = "v1.0";
NullableResponse<TemplateSpecVersionResource> response = await collection.GetIfExistsAsync(templateSpecVersion);
TemplateSpecVersionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    TemplateSpecVersionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
