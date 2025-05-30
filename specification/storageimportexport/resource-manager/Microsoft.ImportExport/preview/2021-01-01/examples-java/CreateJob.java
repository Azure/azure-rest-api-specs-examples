
import com.azure.resourcemanager.storageimportexport.models.DriveStatus;
import com.azure.resourcemanager.storageimportexport.models.JobDetails;
import com.azure.resourcemanager.storageimportexport.models.ReturnAddress;
import com.azure.resourcemanager.storageimportexport.models.ReturnShipping;
import java.util.Arrays;

/**
 * Samples for Jobs Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/CreateJob.
     * json
     */
    /**
     * Sample code: Create import job.
     * 
     * @param manager Entry point to StorageImportExportManager.
     */
    public static void
        createImportJob(com.azure.resourcemanager.storageimportexport.StorageImportExportManager manager) {
        manager.jobs().define("myJob").withExistingResourceGroup("myResourceGroup").withRegion("West US")
            .withProperties(new JobDetails().withStorageAccountId(
                "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.ClassicStorage/storageAccounts/test")
                .withJobType("Import")
                .withReturnAddress(new ReturnAddress().withRecipientName("Test").withStreetAddress1("Street1")
                    .withStreetAddress2("street2").withCity("Redmond").withStateOrProvince("wa")
                    .withPostalCode("fakeTokenPlaceholder").withCountryOrRegion("USA").withPhone("4250000000")
                    .withEmail("Test@contoso.com"))
                .withReturnShipping(new ReturnShipping().withCarrierName("FedEx").withCarrierAccountNumber("989ffff"))
                .withDiagnosticsPath("waimportexport").withLogLevel("Verbose").withBackupDriveManifest(true)
                .withDriveList(Arrays.asList(new DriveStatus().withDriveId("9CA995BB")
                    .withBitLockerKey("fakeTokenPlaceholder")
                    .withManifestFile("\\8a0c23f7-14b7-470a-9633-fcd46590a1bc.manifest")
                    .withManifestHash("4228EC5D8E048CB9B515338C789314BE8D0B2FDBC7C7A0308E1C826242CDE74E")
                    .withDriveHeaderHash(
                        "0:1048576:FB6B6ED500D49DA6E0D723C98D42C657F2881CC13357C28DCECA6A524F1292501571A321238540E621AB5BD9C9A32637615919A75593E6CB5C1515DAE341CABF;135266304:143360:C957A189AFC38C4E80731252301EB91427CE55E61448FA3C73C6FDDE70ABBC197947EC8D0249A2C639BB10B95957D5820A4BE8DFBBF76FFFA688AE5CE0D42EC3"))))
            .create();
    }
}
