using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedServices;
using Azure.ResourceManager.ManagedServices.Models;

// Generated from example definition: specification/managedservices/resource-manager/Microsoft.ManagedServices/stable/2022-10-01/examples/DeleteRegistrationDefinition.json
// this example is just showing the usage of "RegistrationDefinitions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedServicesRegistrationResource created on azure
// for more information of creating ManagedServicesRegistrationResource, please refer to the document of ManagedServicesRegistrationResource
string scope = "subscription/0afefe50-734e-4610-8a82-a144ahf49dea";
string registrationId = "26c128c2-fefa-4340-9bb1-6e081c90ada2";
ResourceIdentifier managedServicesRegistrationResourceId = ManagedServicesRegistrationResource.CreateResourceIdentifier(scope, registrationId);
ManagedServicesRegistrationResource managedServicesRegistration = client.GetManagedServicesRegistrationResource(managedServicesRegistrationResourceId);

// invoke the operation
await managedServicesRegistration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
