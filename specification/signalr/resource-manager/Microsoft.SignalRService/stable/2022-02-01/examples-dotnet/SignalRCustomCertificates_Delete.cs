using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SignalR;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomCertificates_Delete.json
// this example is just showing the usage of "SignalRCustomCertificates_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SignalRCustomCertificateResource created on azure
// for more information of creating SignalRCustomCertificateResource, please refer to the document of SignalRCustomCertificateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
string certificateName = "myCert";
ResourceIdentifier signalRCustomCertificateResourceId = SignalRCustomCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, certificateName);
SignalRCustomCertificateResource signalRCustomCertificate = client.GetSignalRCustomCertificateResource(signalRCustomCertificateResourceId);

// invoke the operation
await signalRCustomCertificate.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
