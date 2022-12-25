using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedServices.Models;
using Azure.ResourceManager.ManagedServices;

// Generated from example definition: specification/managedservices/resource-manager/Microsoft.ManagedServices/stable/2022-10-01/examples/GetRegistrationDefinition.json
// this example is just showing the usage of "RegistrationDefinitions_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this ManagedServicesRegistrationResource
string scope = "subscription/0afefe50-734e-4610-8a82-a144ahf49dea";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
ManagedServicesRegistrationCollection collection = client.GetManagedServicesRegistrations(scopeId);

// invoke the operation
string registrationId = "26c128c2-fefa-4340-9bb1-6e081c90ada2";
bool result = await collection.ExistsAsync(registrationId);

Console.WriteLine($"Succeeded: {result}");
