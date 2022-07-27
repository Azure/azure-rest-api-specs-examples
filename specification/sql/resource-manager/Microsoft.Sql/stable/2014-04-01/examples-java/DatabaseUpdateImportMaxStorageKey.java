import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.AuthenticationType;
import com.azure.resourcemanager.sql.models.ExtensionName;
import com.azure.resourcemanager.sql.models.ImportExtensionRequest;
import com.azure.resourcemanager.sql.models.ImportOperationMode;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/** Samples for Databases CreateImportOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseUpdateImportMaxStorageKey.json
     */
    /**
     * Sample code: Import bacpac into an existing database Max with storage key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importBacpacIntoAnExistingDatabaseMaxWithStorageKey(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .createImportOperation(
                "sqlcrudtest-4799",
                "sqlcrudtest-5961",
                "testdb",
                ExtensionName.IMPORT,
                new ImportExtensionRequest()
                    .withName("Import")
                    .withType("Microsoft.Sql/servers/databases/extensions")
                    .withOperationMode(ImportOperationMode.IMPORT)
                    .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                    .withStorageKey(
                        "sdlfkjdsf+sdlfkjsdlkfsjdfLDKFJSDLKFDFKLjsdfksjdflsdkfD2342309432849328479324/3RSD==")
                    .withStorageUri("https://test.blob.core.windows.net/bacpacs/testbacpac.bacpac")
                    .withAdministratorLogin("dummyLogin")
                    .withAdministratorLoginPassword("Un53cuRE!")
                    .withAuthenticationType(AuthenticationType.SQL),
                Context.NONE);
    }
}
