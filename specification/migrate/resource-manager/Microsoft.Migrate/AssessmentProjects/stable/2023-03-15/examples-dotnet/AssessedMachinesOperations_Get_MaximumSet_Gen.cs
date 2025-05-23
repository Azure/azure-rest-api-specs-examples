using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessedMachinesOperations_Get_MaximumSet_Gen.json
// this example is just showing the usage of "AssessedMachinesOperations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessedMachineResource created on azure
// for more information of creating MigrationAssessedMachineResource, please refer to the document of MigrationAssessedMachineResource
string subscriptionId = "D8E1C413-E65F-40C0-8A7E-743D6B7A6AE9";
string resourceGroupName = "rgopenapi";
string projectName = "pavqtntysjn";
string groupName = "smawqdmhfngray";
string assessmentName = "qjlumxyqsitd";
string assessedMachineName = "oqxjeheiipjmuo";
ResourceIdentifier migrationAssessedMachineResourceId = MigrationAssessedMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, groupName, assessmentName, assessedMachineName);
MigrationAssessedMachineResource migrationAssessedMachine = client.GetMigrationAssessedMachineResource(migrationAssessedMachineResourceId);

// invoke the operation
MigrationAssessedMachineResource result = await migrationAssessedMachine.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationAssessedMachineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
