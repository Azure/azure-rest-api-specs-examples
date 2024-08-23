
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/Operations_List.json
     */
    /**
     * Sample code: Lists all of the available RP operations.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void
        listsAllOfTheAvailableRPOperations(com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
