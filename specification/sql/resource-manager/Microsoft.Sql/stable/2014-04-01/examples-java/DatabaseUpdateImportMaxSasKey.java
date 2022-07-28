import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.AuthenticationType;
import com.azure.resourcemanager.sql.models.ExtensionName;
import com.azure.resourcemanager.sql.models.ImportExtensionRequest;
import com.azure.resourcemanager.sql.models.ImportOperationMode;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/** Samples for Databases CreateImportOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseUpdateImportMaxSasKey.json
     */
    /**
     * Sample code: Import bacpac into an existing database Max with SAS key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importBacpacIntoAnExistingDatabaseMaxWithSASKey(
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
                    .withStorageKeyType(StorageKeyType.SHARED_ACCESS_KEY)
                    .withStorageKey(
                        "?sr=b&sp=rw&se=2018-01-01T00%3A00%3A00Z&sig=sdfsdfklsdjflSLIFJLSIEJFLKSDJFDd/%2wdfskdjf3%3D&sv=2015-07-08")
                    .withStorageUri("https://test.blob.core.windows.net/bacpacs/testbacpac.bacpac")
                    .withAdministratorLogin("dummyLogin")
                    .withAdministratorLoginPassword("Un53cuRE!")
                    .withAuthenticationType(AuthenticationType.SQL),
                Context.NONE);
    }
}
