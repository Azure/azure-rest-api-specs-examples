import com.azure.core.util.Context;

/** Samples for Operations GetLocationHeaderResult. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetLocationHeader.json
     */
    /**
     * Sample code: Get location header result.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getLocationHeaderResult(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .operations()
            .getLocationHeaderResultWithResponse(
                "resourceGroup1", "workspace1", "01234567-89ab-4def-0123-456789abcdef", Context.NONE);
    }
}
