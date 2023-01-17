/** Samples for CustomImages Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/CustomImages_Delete.json
     */
    /**
     * Sample code: CustomImages_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void customImagesDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .customImages()
            .delete("resourceGroupName", "{labName}", "{customImageName}", com.azure.core.util.Context.NONE);
    }
}
