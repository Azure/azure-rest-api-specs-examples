
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_Maximum.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void operationsListMaximum(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
