using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Reservations.Models;
using Azure.ResourceManager.Reservations;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/CalculateRefund.json
// this example is just showing the usage of "CalculateRefund_Post" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReservationOrderResource created on azure
// for more information of creating ReservationOrderResource, please refer to the document of ReservationOrderResource
Guid reservationOrderId = Guid.Parse("276e7ae4-84d0-4da6-ab4b-d6b94f3557da");
ResourceIdentifier reservationOrderResourceId = ReservationOrderResource.CreateResourceIdentifier(reservationOrderId);
ReservationOrderResource reservationOrder = client.GetReservationOrderResource(reservationOrderResourceId);

// invoke the operation
ReservationCalculateRefundContent content = new ReservationCalculateRefundContent
{
    Id = "/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004",
    Properties = new ReservationCalculateRefundRequestProperties
    {
        Scope = "Reservation",
        ReservationToReturn = new ReservationToReturn
        {
            ReservationId = new ResourceIdentifier("/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004/reservations/40000000-aaaa-bbbb-cccc-100000000000"),
            Quantity = 1,
        },
    },
};
ReservationCalculateRefundResult result = await reservationOrder.CalculateRefundAsync(content);

Console.WriteLine($"Succeeded: {result}");
