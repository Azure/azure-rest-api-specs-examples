using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedServices.Models;
using Azure.ResourceManager.ManagedServices;

// Generated from example definition: specification/managedservices/resource-manager/Microsoft.ManagedServices/stable/2022-10-01/examples/GetRegistrationDefinition.json
// this example is just showing the usage of "RegistrationDefinitions_Get" operation, for the dependent resources, they will have to be created separately.

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
ManagedServicesRegistrationResource result = await managedServicesRegistration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedServicesRegistrationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
