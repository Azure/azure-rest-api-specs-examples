using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedServices;
using Azure.ResourceManager.ManagedServices.Models;

// Generated from example definition: specification/managedservices/resource-manager/Microsoft.ManagedServices/stable/2022-10-01/examples/DeleteRegistrationAssignment.json
// this example is just showing the usage of "RegistrationAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedServicesRegistrationAssignmentResource created on azure
// for more information of creating ManagedServicesRegistrationAssignmentResource, please refer to the document of ManagedServicesRegistrationAssignmentResource
string scope = "subscription/0afefe50-734e-4610-8a82-a144ahf49dea";
string registrationAssignmentId = "26c128c2-fefa-4340-9bb1-6e081c90ada2";
ResourceIdentifier managedServicesRegistrationAssignmentResourceId = ManagedServicesRegistrationAssignmentResource.CreateResourceIdentifier(scope, registrationAssignmentId);
ManagedServicesRegistrationAssignmentResource managedServicesRegistrationAssignment = client.GetManagedServicesRegistrationAssignmentResource(managedServicesRegistrationAssignmentResourceId);

// invoke the operation
await managedServicesRegistrationAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
