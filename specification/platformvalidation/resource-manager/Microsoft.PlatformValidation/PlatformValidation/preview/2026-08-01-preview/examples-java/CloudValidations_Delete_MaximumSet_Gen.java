
/**
 * Samples for CloudValidations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/CloudValidations_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudValidations_Delete_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void cloudValidationsDeleteMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.cloudValidations().delete("rgvalidate", "cvtest01", com.azure.core.util.Context.NONE);
    }
}
