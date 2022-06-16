import com.azure.resourcemanager.iotcentral.models.AppSku;
import com.azure.resourcemanager.iotcentral.models.AppSkuInfo;
import com.azure.resourcemanager.iotcentral.models.SystemAssignedServiceIdentity;
import com.azure.resourcemanager.iotcentral.models.SystemAssignedServiceIdentityType;

/** Samples for Apps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_CreateOrUpdate.json
     */
    /**
     * Sample code: Apps_CreateOrUpdate.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsCreateOrUpdate(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .apps()
            .define("myIoTCentralApp")
            .withRegion("westus")
            .withExistingResourceGroup("resRg")
            .withSku(new AppSkuInfo().withName(AppSku.ST2))
            .withIdentity(
                new SystemAssignedServiceIdentity().withType(SystemAssignedServiceIdentityType.SYSTEM_ASSIGNED))
            .withDisplayName("My IoT Central App")
            .withSubdomain("my-iot-central-app")
            .withTemplate("iotc-pnp-preview@1.0.0")
            .create();
    }
}
