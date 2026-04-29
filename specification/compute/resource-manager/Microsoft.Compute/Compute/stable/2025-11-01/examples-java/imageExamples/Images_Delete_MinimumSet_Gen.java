
/**
 * Samples for Images Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/imageExamples/Images_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Image_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void imageDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getImages().delete("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
