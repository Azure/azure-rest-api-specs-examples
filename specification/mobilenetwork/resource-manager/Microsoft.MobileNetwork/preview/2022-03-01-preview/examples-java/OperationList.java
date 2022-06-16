import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/OperationList.json
     */
    /**
     * Sample code: Get Registration Operations.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getRegistrationOperations(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.operations().list(Context.NONE);
    }
}
