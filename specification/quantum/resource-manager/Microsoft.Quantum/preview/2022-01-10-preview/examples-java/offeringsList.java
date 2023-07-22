/** Samples for Offerings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/offeringsList.json
     */
    /**
     * Sample code: OfferingsList.
     *
     * @param manager Entry point to AzureQuantumManager.
     */
    public static void offeringsList(com.azure.resourcemanager.quantum.AzureQuantumManager manager) {
        manager.offerings().list("westus2", com.azure.core.util.Context.NONE);
    }
}
