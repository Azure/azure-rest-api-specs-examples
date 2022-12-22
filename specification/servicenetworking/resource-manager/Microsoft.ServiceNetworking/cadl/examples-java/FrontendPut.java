import com.azure.resourcemanager.servicenetworking.models.FrontendIpAddressVersion;
import com.azure.resourcemanager.servicenetworking.models.FrontendMode;
import com.azure.resourcemanager.servicenetworking.models.FrontendPropertiesIpAddress;

/** Samples for FrontendsInterface CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendPut.json
     */
    /**
     * Sample code: Put Frontend.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void putFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager
            .frontendsInterfaces()
            .define("publicIp1")
            .withRegion("West US")
            .withExistingTrafficController("rg1", "TC1")
            .withMode(FrontendMode.PUBLIC)
            .withIpAddressVersion(FrontendIpAddressVersion.IPV4)
            .withPublicIpAddress(new FrontendPropertiesIpAddress().withId("resourceUriAsString"))
            .create();
    }
}
