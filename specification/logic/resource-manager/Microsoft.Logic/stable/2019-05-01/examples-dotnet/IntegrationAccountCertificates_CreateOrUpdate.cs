using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_CreateOrUpdate.json
// this example is just showing the usage of "IntegrationAccountCertificates_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationAccountCertificateResource created on azure
// for more information of creating IntegrationAccountCertificateResource, please refer to the document of IntegrationAccountCertificateResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string integrationAccountName = "testIntegrationAccount";
string certificateName = "testCertificate";
ResourceIdentifier integrationAccountCertificateResourceId = IntegrationAccountCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, integrationAccountName, certificateName);
IntegrationAccountCertificateResource integrationAccountCertificate = client.GetIntegrationAccountCertificateResource(integrationAccountCertificateResourceId);

// invoke the operation
IntegrationAccountCertificateData data = new IntegrationAccountCertificateData(new AzureLocation("brazilsouth"))
{
    Key = new IntegrationAccountKeyVaultKeyReference("<keyName>")
    {
        KeyVersion = "87d9764197604449b9b8eb7bd8710868",
        ResourceId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testResourceGroup/providers/microsoft.keyvault/vaults/<keyVaultName>"),
    },
    PublicCertificate = BinaryData.FromString("\"<publicCertificateValue>\""),
};
ArmOperation<IntegrationAccountCertificateResource> lro = await integrationAccountCertificate.UpdateAsync(WaitUntil.Completed, data);
IntegrationAccountCertificateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IntegrationAccountCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
