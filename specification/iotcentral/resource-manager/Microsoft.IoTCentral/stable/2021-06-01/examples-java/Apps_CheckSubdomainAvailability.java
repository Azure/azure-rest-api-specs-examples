import com.azure.core.util.Context;
import com.azure.resourcemanager.iotcentral.models.OperationInputs;

/** Samples for Apps CheckSubdomainAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_CheckSubdomainAvailability.json
     */
    /**
     * Sample code: Apps_SubdomainAvailability.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsSubdomainAvailability(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .apps()
            .checkSubdomainAvailabilityWithResponse(
                new OperationInputs().withName("myiotcentralapp").withType("IoTApps"), Context.NONE);
    }
}
