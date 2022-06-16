import com.azure.core.util.Context;
import com.azure.resourcemanager.iotcentral.models.App;
import com.azure.resourcemanager.iotcentral.models.SystemAssignedServiceIdentity;
import com.azure.resourcemanager.iotcentral.models.SystemAssignedServiceIdentityType;

/** Samples for Apps Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Update.json
     */
    /**
     * Sample code: Apps_Update.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsUpdate(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        App resource =
            manager.apps().getByResourceGroupWithResponse("resRg", "myIoTCentralApp", Context.NONE).getValue();
        resource
            .update()
            .withIdentity(
                new SystemAssignedServiceIdentity().withType(SystemAssignedServiceIdentityType.SYSTEM_ASSIGNED))
            .withDisplayName("My IoT Central App 2")
            .apply();
    }
}
