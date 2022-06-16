import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/OperationDescription_List.json
     */
    /**
     * Sample code: OperationDescription_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void operationDescriptionList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.operations().list(Context.NONE);
    }
}
