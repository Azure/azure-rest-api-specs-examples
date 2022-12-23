import com.azure.core.util.Context;
import com.azure.resourcemanager.servicenetworking.models.Frontend;
import com.azure.resourcemanager.servicenetworking.models.FrontendIpAddressVersion;
import com.azure.resourcemanager.servicenetworking.models.FrontendMode;
import com.azure.resourcemanager.servicenetworking.models.FrontendPropertiesIpAddress;
import com.azure.resourcemanager.servicenetworking.models.FrontendUpdateProperties;

/** Samples for FrontendsInterface Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendPatch.json
     */
    /**
     * Sample code: Update Frontend.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void updateFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        Frontend resource =
            manager.frontendsInterfaces().getWithResponse("rg1", "TC1", "publicIp1", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                new FrontendUpdateProperties()
                    .withMode(FrontendMode.PUBLIC)
                    .withIpAddressVersion(FrontendIpAddressVersion.IPV4)
                    .withPublicIpAddress(new FrontendPropertiesIpAddress().withId("resourceUriAsString")))
            .apply();
    }
}
