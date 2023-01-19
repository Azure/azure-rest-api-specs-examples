using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;
using Azure.ResourceManager.Reservations.Models;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/ChangeDirectoryReservationOrder.json
// this example is just showing the usage of "ReservationOrder_ChangeDirectory" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReservationOrderResource created on azure
// for more information of creating ReservationOrderResource, please refer to the document of ReservationOrderResource
Guid reservationOrderId = Guid.Parse("a075419f-44cc-497f-b68a-14ee811d48b9");
ResourceIdentifier reservationOrderResourceId = ReservationOrderResource.CreateResourceIdentifier(reservationOrderId);
ReservationOrderResource reservationOrder = client.GetReservationOrderResource(reservationOrderResourceId);

// invoke the operation
ChangeDirectoryContent content = new ChangeDirectoryContent()
{
    DestinationTenantId = Guid.Parse("906655ea-30be-4587-9d12-b50e077b0f32"),
};
ChangeDirectoryDetail result = await reservationOrder.ChangeDirectoryAsync(content);

Console.WriteLine($"Succeeded: {result}");
