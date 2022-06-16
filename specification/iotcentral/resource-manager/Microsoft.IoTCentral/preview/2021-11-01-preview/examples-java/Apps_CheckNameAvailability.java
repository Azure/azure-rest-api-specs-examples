import com.azure.core.util.Context;
import com.azure.resourcemanager.iotcentral.models.OperationInputs;

/** Samples for Apps CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_CheckNameAvailability.json
     */
    /**
     * Sample code: Apps_CheckNameAvailability.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsCheckNameAvailability(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .apps()
            .checkNameAvailabilityWithResponse(
                new OperationInputs().withName("myiotcentralapp").withType("IoTApps"), Context.NONE);
    }
}
