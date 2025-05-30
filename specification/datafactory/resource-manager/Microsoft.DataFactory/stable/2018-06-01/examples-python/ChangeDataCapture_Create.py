from azure.identity import DefaultAzureCredential

from azure.mgmt.datafactory import DataFactoryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datafactory
# USAGE
    python change_data_capture_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataFactoryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-1234-12345678abc",
    )

    response = client.change_data_capture.create_or_update(
        resource_group_name="exampleResourceGroup",
        factory_name="exampleFactoryName",
        change_data_capture_name="exampleChangeDataCapture",
        change_data_capture={
            "properties": {
                "Policy": {"mode": "Microbatch", "recurrence": {"frequency": "Minute", "interval": 15}},
                "SourceConnectionsInfo": [
                    {
                        "Connection": {
                            "commonDslConnectorProperties": [
                                {"name": "allowSchemaDrift", "value": True},
                                {"name": "inferDriftedColumnTypes", "value": True},
                                {"name": "format", "value": "delimited"},
                                {
                                    "name": "dateFormats",
                                    "value": [
                                        "MM/dd/yyyy",
                                        "dd/MM/yyyy",
                                        "yyyy/MM/dd",
                                        "MM-dd-yyyy",
                                        "dd-MM-yyyy",
                                        "yyyy-MM-dd",
                                        "dd.MM.yyyy",
                                        "MM.dd.yyyy",
                                        "yyyy.MM.dd",
                                    ],
                                },
                                {
                                    "name": "timestampFormats",
                                    "value": [
                                        "yyyyMMddHHmm",
                                        "yyyyMMdd HHmm",
                                        "yyyyMMddHHmmss",
                                        "yyyyMMdd HHmmss",
                                        "dd-MM-yyyy HH:mm:ss",
                                        "dd-MM-yyyy HH:mm",
                                        "yyyy-M-d H:m:s",
                                        "yyyy-MM-dd\\'T\\'HH:mm:ss\\'Z\\'",
                                        "yyyy-M-d\\'T\\'H:m:s\\'Z\\'",
                                        "yyyy-M-d\\'T\\'H:m:s",
                                        "yyyy-MM-dd\\'T\\'HH:mm:ss",
                                        "yyyy-MM-dd HH:mm:ss",
                                        "yyyy-MM-dd HH:mm",
                                        "yyyy.MM.dd HH:mm:ss",
                                        "MM/dd/yyyy HH:mm:ss",
                                        "M/d/yyyy H:m:s",
                                        "yyyy/MM/dd HH:mm:ss",
                                        "yyyy/M/d H:m:s",
                                        "dd MMM yyyy HH:mm:ss",
                                        "dd MMMM yyyy HH:mm:ss",
                                        "d MMM yyyy H:m:s",
                                        "d MMMM yyyy H:m:s",
                                        "d-M-yyyy H:m:s",
                                        "d-M-yyyy H:m",
                                        "yyyy-M-d H:m",
                                        "MM/dd/yyyy HH:mm",
                                        "M/d/yyyy H:m",
                                        "yyyy/MM/dd HH:mm",
                                        "yyyy/M/d H:m",
                                        "dd MMMM yyyy HH:mm",
                                        "dd MMM yyyy HH:mm",
                                        "d MMMM yyyy H:m",
                                        "d MMM yyyy H:m",
                                        "MM-dd-yyyy hh:mm:ss a",
                                        "MM-dd-yyyy HH:mm:ss",
                                        "MM/dd/yyyy hh:mm:ss a",
                                        "yyyy.MM.dd hh:mm:ss a",
                                        "MM/dd/yyyy",
                                        "dd/MM/yyyy",
                                        "yyyy/MM/dd",
                                        "MM-dd-yyyy",
                                        "dd-MM-yyyy",
                                        "yyyy-MM-dd",
                                        "dd.MM.yyyy",
                                        "MM.dd.yyyy",
                                        "yyyy.MM.dd",
                                    ],
                                },
                                {"name": "enableCdc", "value": True},
                                {"name": "skipInitialLoad", "value": True},
                                {"name": "columnNamesAsHeader", "value": True},
                                {"name": "columnDelimiter", "value": ","},
                                {"name": "escapeChar", "value": "\\\\"},
                                {"name": "quoteChar", "value": '\\"'},
                            ],
                            "isInlineDataset": True,
                            "linkedService": {"referenceName": "amjaAdls03", "type": "LinkedServiceReference"},
                            "linkedServiceType": "AzureBlobFS",
                            "type": "linkedservicetype",
                        },
                        "SourceEntities": [
                            {
                                "name": "source/customer",
                                "properties": {
                                    "dslConnectorProperties": [
                                        {"name": "container", "value": "source"},
                                        {"name": "fileSystem", "value": "source"},
                                        {"name": "folderPath", "value": "customer"},
                                        {"name": "allowSchemaDrift", "value": False},
                                        {"name": "inferDriftedColumnTypes", "value": False},
                                    ],
                                    "schema": [
                                        {"dataType": "short", "name": "CustId"},
                                        {"dataType": "string", "name": "CustName"},
                                        {"dataType": "string", "name": "CustAddres"},
                                        {"dataType": "string", "name": "CustDepName"},
                                        {"dataType": "string", "name": "CustDepLoc"},
                                    ],
                                },
                            },
                            {
                                "name": "source/employee",
                                "properties": {
                                    "dslConnectorProperties": [
                                        {"name": "container", "value": "source"},
                                        {"name": "fileSystem", "value": "source"},
                                        {"name": "folderPath", "value": "employee"},
                                    ],
                                    "schema": [],
                                },
                            },
                            {
                                "name": "lookup",
                                "properties": {
                                    "dslConnectorProperties": [
                                        {"name": "container", "value": "lookup"},
                                        {"name": "fileSystem", "value": "lookup"},
                                        {"name": "allowSchemaDrift", "value": False},
                                        {"name": "inferDriftedColumnTypes", "value": False},
                                    ],
                                    "schema": [
                                        {"dataType": "short", "name": "EmpId"},
                                        {"dataType": "string", "name": "EmpName"},
                                        {"dataType": "string", "name": "HomeAddress"},
                                        {"dataType": "string", "name": "OfficeAddress"},
                                        {"dataType": "integer", "name": "EmpPhoneNumber"},
                                        {"dataType": "string", "name": "DepName"},
                                        {"dataType": "string", "name": "DepLoc"},
                                        {"dataType": "double", "name": "DecimalCol"},
                                    ],
                                },
                            },
                            {
                                "name": "source/justSchema",
                                "properties": {
                                    "dslConnectorProperties": [
                                        {"name": "container", "value": "source"},
                                        {"name": "fileSystem", "value": "source"},
                                        {"name": "folderPath", "value": "justSchema"},
                                        {"name": "allowSchemaDrift", "value": False},
                                        {"name": "inferDriftedColumnTypes", "value": False},
                                    ],
                                    "schema": [
                                        {"dataType": "string", "name": "CustId"},
                                        {"dataType": "string", "name": "CustName"},
                                        {"dataType": "string", "name": "CustAddres"},
                                        {"dataType": "string", "name": "CustDepName"},
                                        {"dataType": "string", "name": "CustDepLoc"},
                                    ],
                                },
                            },
                        ],
                    }
                ],
                "TargetConnectionsInfo": [
                    {
                        "Connection": {
                            "commonDslConnectorProperties": [
                                {"name": "allowSchemaDrift", "value": True},
                                {"name": "inferDriftedColumnTypes", "value": True},
                                {"name": "format", "value": "table"},
                                {"name": "store", "value": "sqlserver"},
                                {"name": "databaseType", "value": "databaseType"},
                                {"name": "database", "value": "database"},
                                {"name": "deletable", "value": False},
                                {"name": "insertable", "value": True},
                                {"name": "updateable", "value": False},
                                {"name": "upsertable", "value": False},
                                {"name": "skipDuplicateMapInputs", "value": True},
                                {"name": "skipDuplicateMapOutputs", "value": True},
                            ],
                            "isInlineDataset": True,
                            "linkedService": {"referenceName": "amjaSql", "type": "LinkedServiceReference"},
                            "linkedServiceType": "AzureSqlDatabase",
                            "type": "linkedservicetype",
                        },
                        "DataMapperMappings": [
                            {
                                "attributeMappingInfo": {"attributeMappings": []},
                                "sourceConnectionReference": {
                                    "connectionName": "amjaAdls03",
                                    "type": "linkedservicetype",
                                },
                                "sourceEntityName": "source/customer",
                                "targetEntityName": "dbo.customer",
                            },
                            {
                                "attributeMappingInfo": {
                                    "attributeMappings": [
                                        {
                                            "attributeReferences": [
                                                {
                                                    "entity": "lookup",
                                                    "entityConnectionReference": {
                                                        "connectionName": "amjaAdls03",
                                                        "type": "linkedservicetype",
                                                    },
                                                    "name": "EmpName",
                                                }
                                            ],
                                            "expression": "upper(EmpName)",
                                            "functionName": "upper",
                                            "name": "Name",
                                            "type": "Derived",
                                        },
                                        {
                                            "attributeReference": {
                                                "entity": "lookup",
                                                "entityConnectionReference": {
                                                    "connectionName": "amjaAdls03",
                                                    "type": "linkedservicetype",
                                                },
                                                "name": "EmpId",
                                            },
                                            "functionName": "",
                                            "name": "PersonID",
                                            "type": "Direct",
                                        },
                                    ]
                                },
                                "sourceConnectionReference": {
                                    "connectionName": "amjaAdls03",
                                    "type": "linkedservicetype",
                                },
                                "sourceEntityName": "lookup",
                                "targetEntityName": "dbo.data_source_table",
                            },
                            {
                                "attributeMappingInfo": {"attributeMappings": []},
                                "sourceConnectionReference": {
                                    "connectionName": "amjaAdls03",
                                    "type": "linkedservicetype",
                                },
                                "sourceEntityName": "source/employee",
                                "targetEntityName": "dbo.employee",
                            },
                            {
                                "attributeMappingInfo": {
                                    "attributeMappings": [
                                        {
                                            "attributeReferences": [
                                                {
                                                    "entity": "source/justSchema",
                                                    "entityConnectionReference": {
                                                        "connectionName": "amjaAdls03",
                                                        "type": "linkedservicetype",
                                                    },
                                                    "name": "CustAddres",
                                                }
                                            ],
                                            "expression": "trim(CustAddres)",
                                            "functionName": "trim",
                                            "name": "CustAddres",
                                            "type": "Derived",
                                        },
                                        {
                                            "attributeReference": {
                                                "entity": "source/justSchema",
                                                "entityConnectionReference": {
                                                    "connectionName": "amjaAdls03",
                                                    "type": "linkedservicetype",
                                                },
                                                "name": "CustDepLoc",
                                            },
                                            "name": "CustDepLoc",
                                            "type": "Direct",
                                        },
                                        {
                                            "attributeReferences": [
                                                {
                                                    "entity": "source/justSchema",
                                                    "entityConnectionReference": {
                                                        "connectionName": "amjaAdls03",
                                                        "type": "linkedservicetype",
                                                    },
                                                    "name": "CustName",
                                                },
                                                {
                                                    "entity": "source/justSchema",
                                                    "entityConnectionReference": {
                                                        "connectionName": "amjaAdls03",
                                                        "type": "linkedservicetype",
                                                    },
                                                    "name": "CustDepName",
                                                },
                                            ],
                                            "expression": 'concat(CustName, " -> ", CustDepName)',
                                            "functionName": "",
                                            "name": "CustDepName",
                                            "type": "Derived",
                                        },
                                        {
                                            "attributeReference": {
                                                "entity": "source/justSchema",
                                                "entityConnectionReference": {
                                                    "connectionName": "amjaAdls03",
                                                    "type": "linkedservicetype",
                                                },
                                                "name": "CustId",
                                            },
                                            "functionName": "",
                                            "name": "CustId",
                                            "type": "Direct",
                                        },
                                        {
                                            "attributeReference": {
                                                "entity": "source/justSchema",
                                                "entityConnectionReference": {
                                                    "connectionName": "amjaAdls03",
                                                    "type": "linkedservicetype",
                                                },
                                                "name": "CustName",
                                            },
                                            "name": "CustName",
                                            "type": "Direct",
                                        },
                                    ]
                                },
                                "sourceConnectionReference": {
                                    "connectionName": "amjaAdls03",
                                    "type": "linkedservicetype",
                                },
                                "sourceEntityName": "source/justSchema",
                                "targetEntityName": "dbo.justSchema",
                            },
                        ],
                        "Relationships": [],
                        "TargetEntities": [
                            {
                                "name": "dbo.employee",
                                "properties": {
                                    "dslConnectorProperties": [
                                        {"name": "schemaName", "value": "dbo"},
                                        {"name": "tableName", "value": "employee"},
                                    ],
                                    "schema": [],
                                },
                            },
                            {
                                "name": "dbo.justSchema",
                                "properties": {
                                    "dslConnectorProperties": [
                                        {"name": "schemaName", "value": "dbo"},
                                        {"name": "tableName", "value": "justSchema"},
                                        {"name": "allowSchemaDrift", "value": True},
                                        {"name": "inferDriftedColumnTypes", "value": True},
                                    ],
                                    "schema": [],
                                },
                            },
                            {
                                "name": "dbo.customer",
                                "properties": {
                                    "dslConnectorProperties": [
                                        {"name": "schemaName", "value": "dbo"},
                                        {"name": "tableName", "value": "customer"},
                                        {"name": "allowSchemaDrift", "value": False},
                                        {"name": "inferDriftedColumnTypes", "value": False},
                                    ],
                                    "schema": [
                                        {"dataType": "integer", "name": "CustId"},
                                        {"dataType": "string", "name": "CustName"},
                                        {"dataType": "string", "name": "CustAddres"},
                                        {"dataType": "string", "name": "CustDeptName"},
                                        {"dataType": "string", "name": "CustEmail"},
                                    ],
                                },
                            },
                            {
                                "name": "dbo.data_source_table",
                                "properties": {
                                    "dslConnectorProperties": [
                                        {"name": "schemaName", "value": "dbo"},
                                        {"name": "tableName", "value": "data_source_table"},
                                        {"name": "allowSchemaDrift", "value": False},
                                        {"name": "inferDriftedColumnTypes", "value": False},
                                        {"name": "defaultToUpsert", "value": False},
                                    ],
                                    "schema": [
                                        {"dataType": "integer", "name": "PersonID"},
                                        {"dataType": "string", "name": "Name"},
                                        {"dataType": "timestamp", "name": "LastModifytime"},
                                    ],
                                },
                            },
                        ],
                    }
                ],
                "allowVNetOverride": False,
                "description": "Sample demo change data capture to transfer data from delimited (csv) to Azure SQL Database with automapped and non-automapped mappings.",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ChangeDataCapture_Create.json
if __name__ == "__main__":
    main()
