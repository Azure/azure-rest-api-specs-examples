import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void operationsList(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.operations().list(Context.NONE);
    }
}
