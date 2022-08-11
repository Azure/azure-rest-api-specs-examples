import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listOperations.json
     */
    /**
     * Sample code: Lists all of the available Automanage REST API operations.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listsAllOfTheAvailableAutomanageRESTAPIOperations(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.operations().list(Context.NONE);
    }
}
