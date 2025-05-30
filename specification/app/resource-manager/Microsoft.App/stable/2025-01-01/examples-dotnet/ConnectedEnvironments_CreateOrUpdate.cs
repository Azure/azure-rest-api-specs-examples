using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ConnectedEnvironments_CreateOrUpdate.json
// this example is just showing the usage of "ConnectedEnvironments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerAppConnectedEnvironmentResource
ContainerAppConnectedEnvironmentCollection collection = resourceGroupResource.GetContainerAppConnectedEnvironments();

// invoke the operation
string connectedEnvironmentName = "testenv";
ContainerAppConnectedEnvironmentData data = new ContainerAppConnectedEnvironmentData(new AzureLocation("East US"))
{
    StaticIP = IPAddress.Parse("1.2.3.4"),
    DaprAIConnectionString = "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/",
    CustomDomainConfiguration = new ContainerAppCustomDomainConfiguration
    {
        DnsSuffix = "www.my-name.com",
        CertificateValue = Encoding.UTF8.GetBytes("Y2VydA=="),
        CertificatePassword = "private key password",
    },
};
ArmOperation<ContainerAppConnectedEnvironmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, connectedEnvironmentName, data);
ContainerAppConnectedEnvironmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppConnectedEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
