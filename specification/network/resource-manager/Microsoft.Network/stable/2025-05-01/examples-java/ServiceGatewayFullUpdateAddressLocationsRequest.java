
import com.azure.resourcemanager.network.models.AddressUpdateAction;
import com.azure.resourcemanager.network.models.ServiceGatewayAddress;
import com.azure.resourcemanager.network.models.ServiceGatewayAddressLocation;
import com.azure.resourcemanager.network.models.ServiceGatewayUpdateAddressLocationsRequest;
import com.azure.resourcemanager.network.models.UpdateAction;
import java.util.Arrays;

/**
 * Samples for ServiceGateways UpdateAddressLocations.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ServiceGatewayFullUpdateAddressLocationsRequest.json
     */
    /**
     * Sample code: Full Update: Create, update, or delete address locations in the service gateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void fullUpdateCreateUpdateOrDeleteAddressLocationsInTheServiceGateway(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceGateways().updateAddressLocations("rg1", "sg",
            new ServiceGatewayUpdateAddressLocationsRequest().withAction(UpdateAction.FULL_UPDATE)
                .withAddressLocations(Arrays.asList(
                    new ServiceGatewayAddressLocation()
                        .withAddressLocation("192.0.0.1").withAddressUpdateAction(AddressUpdateAction.FULL_UPDATE)
                        .withAddresses(
                            Arrays.asList(new ServiceGatewayAddress()
                                .withAddress("10.0.0.4").withServices(Arrays.asList("Service1")))),
                    new ServiceGatewayAddressLocation().withAddressLocation("192.0.0.2")
                        .withAddressUpdateAction(AddressUpdateAction.PARTIAL_UPDATE)
                        .withAddresses(Arrays.asList(
                            new ServiceGatewayAddress().withAddress("10.0.0.5").withServices(Arrays.asList("Service2")),
                            new ServiceGatewayAddress().withAddress("10.0.0.6"))))),
            com.azure.core.util.Context.NONE);
    }
}
