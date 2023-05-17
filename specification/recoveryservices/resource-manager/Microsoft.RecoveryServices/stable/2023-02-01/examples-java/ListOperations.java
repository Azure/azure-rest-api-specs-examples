/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void listOperations(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
