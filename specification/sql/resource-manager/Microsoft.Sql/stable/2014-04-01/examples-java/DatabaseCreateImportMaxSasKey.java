import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.AuthenticationType;
import com.azure.resourcemanager.sql.models.DatabaseEdition;
import com.azure.resourcemanager.sql.models.ImportRequest;
import com.azure.resourcemanager.sql.models.ServiceObjectiveName;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/** Samples for Databases ImportMethod. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseCreateImportMaxSasKey.json
     */
    /**
     * Sample code: Import bacpac into new database Max with SAS key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importBacpacIntoNewDatabaseMaxWithSASKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .importMethod(
                "sqlcrudtest-4799",
                "sqlcrudtest-5961",
                new ImportRequest()
                    .withStorageKeyType(StorageKeyType.SHARED_ACCESS_KEY)
                    .withStorageKey(
                        "?sr=b&sp=rw&se=2018-01-01T00%3A00%3A00Z&sig=sdfsdfklsdjflSLIFJLSIEJFLKSDJFDd/%2wdfskdjf3%3D&sv=2015-07-08")
                    .withStorageUri("https://test.blob.core.windows.net/bacpacs/testbacpac.bacpac")
                    .withAdministratorLogin("dummyLogin")
                    .withAdministratorLoginPassword("Un53cuRE!")
                    .withAuthenticationType(AuthenticationType.SQL)
                    .withDatabaseName("TestDbImport")
                    .withEdition(DatabaseEdition.BASIC)
                    .withServiceObjectiveName(ServiceObjectiveName.BASIC)
                    .withMaxSizeBytes("2147483648"),
                Context.NONE);
    }
}
