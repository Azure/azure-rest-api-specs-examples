import com.azure.resourcemanager.azurestackhci.models.ExtendedLocation;
import com.azure.resourcemanager.azurestackhci.models.ExtendedLocationTypes;
import com.azure.resourcemanager.azurestackhci.models.IpConfiguration;
import com.azure.resourcemanager.azurestackhci.models.IpConfigurationProperties;
import com.azure.resourcemanager.azurestackhci.models.IpConfigurationPropertiesSubnet;
import java.util.Arrays;

/** Samples for NetworkInterfacesOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutNetworkInterface.json
     */
    /**
     * Sample code: PutNetworkInterface.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void putNetworkInterface(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .networkInterfacesOperations()
            .define("test-nic")
            .withRegion("West US2")
            .withExistingResourceGroup("test-rg")
            .withExtendedLocation(
                new ExtendedLocation()
                    .withName(
                        "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location")
                    .withType(ExtendedLocationTypes.CUSTOM_LOCATION))
            .withIpConfigurations(
                Arrays
                    .asList(
                        new IpConfiguration()
                            .withName("ipconfig-sample")
                            .withProperties(
                                new IpConfigurationProperties()
                                    .withSubnet(new IpConfigurationPropertiesSubnet().withId("test-lnet")))))
            .create();
    }
}
