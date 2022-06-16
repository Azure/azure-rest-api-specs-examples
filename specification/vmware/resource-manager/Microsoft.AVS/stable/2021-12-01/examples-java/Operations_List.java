import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void operationsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.operations().list(Context.NONE);
    }
}
