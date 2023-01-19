import com.azure.resourcemanager.storageimportexport.models.Export;
import com.azure.resourcemanager.storageimportexport.models.JobDetails;
import com.azure.resourcemanager.storageimportexport.models.ReturnAddress;
import com.azure.resourcemanager.storageimportexport.models.ReturnShipping;
import java.util.Arrays;

/** Samples for Jobs Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/CreateExportJob.json
     */
    /**
     * Sample code: Create export job.
     *
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void createExportJob(
        com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager
            .jobs()
            .define("myExportJob")
            .withExistingResourceGroup("myResourceGroup")
            .withRegion("West US")
            .withProperties(
                new JobDetails()
                    .withStorageAccountId(
                        "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.ClassicStorage/storageAccounts/test")
                    .withJobType("Export")
                    .withReturnAddress(
                        new ReturnAddress()
                            .withRecipientName("Test")
                            .withStreetAddress1("Street1")
                            .withStreetAddress2("street2")
                            .withCity("Redmond")
                            .withStateOrProvince("wa")
                            .withPostalCode("fakeTokenPlaceholder")
                            .withCountryOrRegion("USA")
                            .withPhone("4250000000")
                            .withEmail("Test@contoso.com"))
                    .withReturnShipping(
                        new ReturnShipping().withCarrierName("FedEx").withCarrierAccountNumber("989ffff"))
                    .withDiagnosticsPath("waimportexport")
                    .withLogLevel("Verbose")
                    .withBackupDriveManifest(true)
                    .withExport(new Export().withBlobPathPrefix(Arrays.asList("/"))))
            .create();
    }
}
