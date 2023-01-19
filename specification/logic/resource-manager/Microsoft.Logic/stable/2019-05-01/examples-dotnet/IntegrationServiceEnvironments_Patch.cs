using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Patch.json
// this example is just showing the usage of "IntegrationServiceEnvironments_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationServiceEnvironmentResource created on azure
// for more information of creating IntegrationServiceEnvironmentResource, please refer to the document of IntegrationServiceEnvironmentResource
string subscriptionId = "f34b22a3-2202-4fb1-b040-1332bd928c84";
string resourceGroup = "testResourceGroup";
string integrationServiceEnvironmentName = "testIntegrationServiceEnvironment";
ResourceIdentifier integrationServiceEnvironmentResourceId = IntegrationServiceEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroup, integrationServiceEnvironmentName);
IntegrationServiceEnvironmentResource integrationServiceEnvironment = client.GetIntegrationServiceEnvironmentResource(integrationServiceEnvironmentResourceId);

// invoke the operation
IntegrationServiceEnvironmentData data = new IntegrationServiceEnvironmentData(new AzureLocation("placeholder"))
{
    Sku = new IntegrationServiceEnvironmentSku()
    {
        Name = IntegrationServiceEnvironmentSkuName.Developer,
        Capacity = 0,
    },
    Tags =
    {
    ["tag1"] = "value1",
    },
};
ArmOperation<IntegrationServiceEnvironmentResource> lro = await integrationServiceEnvironment.UpdateAsync(WaitUntil.Completed, data);
IntegrationServiceEnvironmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IntegrationServiceEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
