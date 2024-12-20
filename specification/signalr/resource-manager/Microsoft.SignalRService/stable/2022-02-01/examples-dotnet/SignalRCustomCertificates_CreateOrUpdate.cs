using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SignalR;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomCertificates_CreateOrUpdate.json
// this example is just showing the usage of "SignalRCustomCertificates_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SignalRResource created on azure
// for more information of creating SignalRResource, please refer to the document of SignalRResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
ResourceIdentifier signalRResourceId = SignalRResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
SignalRResource signalR = client.GetSignalRResource(signalRResourceId);

// get the collection of this SignalRCustomCertificateResource
SignalRCustomCertificateCollection collection = signalR.GetSignalRCustomCertificates();

// invoke the operation
string certificateName = "myCert";
SignalRCustomCertificateData data = new SignalRCustomCertificateData(new Uri("https://myvault.keyvault.azure.net/"), "mycert")
{
    KeyVaultSecretVersion = "bb6a44b2743f47f68dad0d6cc9756432",
};
ArmOperation<SignalRCustomCertificateResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, certificateName, data);
SignalRCustomCertificateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SignalRCustomCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
