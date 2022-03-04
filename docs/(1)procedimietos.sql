USE [SIM]
GO
/****** Object:  StoredProcedure [dbo].[Utilitarios_Reconstruir_Encabezados]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Utilitarios_Reconstruir_Encabezados]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Utilitarios_Reconstruir_Encabezados]
GO
/****** Object:  StoredProcedure [dbo].[Utilitarios_Reconstruir_Un_Encabezado]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Utilitarios_Reconstruir_Un_Encabezado]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Utilitarios_Reconstruir_Un_Encabezado]
GO
/****** Object:  StoredProcedure [dbo].[Utilitarios_Clasif_Produsctos_Textos]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Utilitarios_Clasif_Produsctos_Textos]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Utilitarios_Clasif_Produsctos_Textos]
GO
/****** Object:  StoredProcedure [dbo].[Reporte1]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Reporte1]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Reporte1]
GO
/****** Object:  StoredProcedure [dbo].[Saldo_Productos]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Saldo_Productos]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Saldo_Productos]
GO
/****** Object:  StoredProcedure [dbo].[Act_Precios_Recepcion]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Act_Precios_Recepcion]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Act_Precios_Recepcion]
GO
/****** Object:  StoredProcedure [dbo].[Utilitarios_BorrarTablas]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Utilitarios_BorrarTablas]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Utilitarios_BorrarTablas]
GO
/****** Object:  StoredProcedure [dbo].[UpdateDetalles]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[UpdateDetalles]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[UpdateDetalles]
GO
/****** Object:  StoredProcedure [dbo].[Total_ImporteCUC_Facturas]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Total_ImporteCUC_Facturas]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Total_ImporteCUC_Facturas]
GO
/****** Object:  StoredProcedure [dbo].[DeleteDetalles]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[DeleteDetalles]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[DeleteDetalles]
GO
/****** Object:  StoredProcedure [dbo].[Recalcular_Recepcion]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Recalcular_Recepcion]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Recalcular_Recepcion]
GO
/****** Object:  StoredProcedure [dbo].[Rep_Conciliaciones_Facturas]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Conciliaciones_Facturas]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Rep_Conciliaciones_Facturas]
GO
/****** Object:  StoredProcedure [dbo].[Rep_Costo_Mensual_Indicadores]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Costo_Mensual_Indicadores]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Rep_Costo_Mensual_Indicadores]
GO
/****** Object:  StoredProcedure [dbo].[Rep_Movimientos_Producto]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Movimientos_Producto]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Rep_Movimientos_Producto]
GO
/****** Object:  StoredProcedure [dbo].[Rep_Comprativo_Inventarios]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Comprativo_Inventarios]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Rep_Comprativo_Inventarios]
GO
/****** Object:  StoredProcedure [dbo].[Rep_Movimiento_Mercancias]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Movimiento_Mercancias]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Rep_Movimiento_Mercancias]
GO
/****** Object:  StoredProcedure [dbo].[Rep_Pedido_A_Fabrica]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Pedido_A_Fabrica]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Rep_Pedido_A_Fabrica]
GO
/****** Object:  StoredProcedure [dbo].[Rep_Rotacion_Mercancias]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Rotacion_Mercancias]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Rep_Rotacion_Mercancias]
GO
/****** Object:  StoredProcedure [dbo].[Saldo_Un_Producto]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Saldo_Un_Producto]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Saldo_Un_Producto]
GO
/****** Object:  StoredProcedure [dbo].[Saldos_Productos_Almacen]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Saldos_Productos_Almacen]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Saldos_Productos_Almacen]
GO
/****** Object:  StoredProcedure [dbo].[Ad_Hoc_Distributed_Queries]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Ad_Hoc_Distributed_Queries]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Ad_Hoc_Distributed_Queries]
GO
/****** Object:  StoredProcedure [dbo].[BackUp_DB]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[BackUp_DB]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[BackUp_DB]
GO
/****** Object:  StoredProcedure [dbo].[Rep_Costo_Mensual]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Costo_Mensual]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Rep_Costo_Mensual]
GO
/****** Object:  StoredProcedure [dbo].[Saldos_Acumulados_Productos_Almacen]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Saldos_Acumulados_Productos_Almacen]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Saldos_Acumulados_Productos_Almacen]
GO
/****** Object:  StoredProcedure [dbo].[Historico_Productos]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Historico_Productos]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Historico_Productos]
GO
/****** Object:  StoredProcedure [dbo].[Restore_DB]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Restore_DB]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Restore_DB]
GO
/****** Object:  StoredProcedure [dbo].[Importar_Saldos_Almacen]    Script Date: 12/04/2013 14:12:46 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Importar_Saldos_Almacen]') AND type in (N'P', N'PC'))
DROP PROCEDURE [dbo].[Importar_Saldos_Almacen]
GO
/****** Object:  StoredProcedure [dbo].[Importar_Saldos_Almacen]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Importar_Saldos_Almacen]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Importar_Saldos_Almacen] @cod_almacen varchar(50), @xls varchar(4000), @nombre_hoja varchar(256) 
AS
BEGIN

	SET NOCOUNT ON;

    exec sp_configure ''Ad Hoc Distributed Queries'', 1
    reconfigure

    declare @cad_sql nvarchar(4000);

    /*
    IF EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N''[dbo].[XLSImportTemp]'') AND type in (N''U''))
      BEGIN
        DROP TABLE [dbo].[XLSImportTemp];
      END
    */
      
    begin try
        DROP TABLE [dbo].[XLSImportTemp];
    end try
    begin catch 
        select * from Nom_Almacenes;
    end catch
    
      

    set @cad_sql = N''SELECT * INTO XLSImportTemp FROM OPENDATASOURCE(''''Microsoft.Jet.OLEDB.4.0'''','' +
                    ''''''Data Source='' + @xls + '';Extended Properties=Excel 8.0'''')...['' + @nombre_hoja + ''$]'';
    execute sp_executesql @cad_sql;

    IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N''[dbo].[XLSImport]'') AND type in (N''U''))
      BEGIN
        CREATE TABLE [dbo].[XLSImport]
        (
          [cod_almacen] [varchar](50) NULL,
	      [cod_producto] [varchar](50) NULL,
	      [saldo] [money] NULL
        )
      END

    delete from XLSImport where cod_almacen = @cod_almacen;

    insert into XLSImport (cod_almacen, cod_producto, saldo)
      select @cod_almacen, codigo, saldo
        from XLSImportTemp;                
        
    update XLSImport
      set cod_producto = SUBSTRING(cod_producto, 5, LEN(cod_producto) - 4)
      where cod_almacen = @cod_almacen;    
                    
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Restore_DB]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Restore_DB]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Restore_DB] @path varchar(256)
AS
BEGIN

  set @path = @path + ''\SIM.bak'';

  RESTORE DATABASE SIM
   FROM DISK = @path;

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Historico_Productos]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Historico_Productos]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Historico_Productos] @cod_producto varchar(50)
AS
BEGIN
	SET NOCOUNT ON;

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Saldos_Acumulados_Productos_Almacen]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Saldos_Acumulados_Productos_Almacen]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Saldos_Acumulados_Productos_Almacen] @cod_almacen varchar(50)
AS
BEGIN

	SET NOCOUNT ON;
	
	truncate table Saldos_Productos_Almacen_Table;
	
    insert into Saldos_Productos_Almacen_Table (cod_producto, descripcion, u_medida, existencia, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta)
      select d.cod_producto, p.descripcion, p.u_medida, SUM(d.cantidad), 0, 0, p.p_venta, 0
        from TransDetalles d inner join Nom_Productos p on d.cod_producto = p.codigo 
        where d.cod_almacen = @cod_almacen and (d.cod_transaccion = ''11'' or cod_transaccion = ''22'') 
        group by d.cod_producto, p.descripcion, p.u_medida, p.p_venta;
    
  update Saldos_Productos_Almacen_Table set saldo_inicial = 0 where saldo_inicial is null;
  update Saldos_Productos_Almacen_Table set entradas = 0 where entradas is null;
  update Saldos_Productos_Almacen_Table set salidas = 0 where salidas is null;
  update Saldos_Productos_Almacen_Table set existencia = 0 where existencia is null;
    
  delete from Saldos_Productos_Almacen_Table where existencia = 0;  

    insert into Saldos_Productos_Almacen_Table (cod_producto, descripcion, u_medida, existencia, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta)
      select h.cod_producto, p.descripcion, p.u_medida, 0, 0, 0, p.p_venta, 0
        from Entradas_Historico h inner join Nom_Productos p on h.cod_producto = p.codigo 
        where h.cod_almacen = @cod_almacen
          and h.cod_producto not in (select cod_producto from Saldos_Productos_Almacen_Table) 
        
/* CALCULO DEL PRECIO DE COSTO PROMEDIO */   
  truncate table no_operaciones;
  insert into no_operaciones (cod_producto, no_operac)    
    select cod_producto, COUNT(*) as no_operac
      from TransDetalles 
      where cod_almacen = @cod_almacen 
        and (cod_transaccion = ''11'' or cod_transaccion = ''22'')
        and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
      group by cod_producto;
  select cod_producto, COUNT(*) as no_operac
    into #noOperac 
    from Entradas_Historico 
    where cod_almacen = @cod_almacen      
      and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
    group by cod_producto;
  insert into No_Operaciones (cod_producto, no_operac)
    select cod_producto, 0
      from #noOperac
      where cod_producto not in (select cod_producto from No_Operaciones)  
  update no
    set no.no_operac = no.no_operac + h.no_operac
    from No_Operaciones no, #noOperac h
    where no.cod_producto = h.cod_producto;     
            
  truncate table costos_promedios;
  insert into costos_promedios (cod_producto, pcosto)    
    select cod_producto, SUM(round(pcosto,2)) as pcosto
      from TransDetalles 
      where cod_almacen = @cod_almacen 
        and (cod_transaccion = ''11'' or cod_transaccion = ''22'')
        and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
      group by cod_producto;
  select cod_producto, SUM(p_costo) as pcosto
    into #pCosto
    from Entradas_Historico 
    where cod_almacen = @cod_almacen    
      and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
    group by cod_producto; 
  insert into Costos_Promedios (cod_producto, pcosto)
    select cod_producto, 0
      from #pCosto  
      where cod_producto not in (select cod_producto from Costos_Promedios)  
  update pc
    set pc.pcosto = pc.pcosto + h.pcosto
    from Costos_Promedios pc, #pCosto h
    where pc.cod_producto = h.cod_producto;

  select cod_producto, SUM(cantidad) as cantidad
    into #cAcumulada
    from Entradas_Historico
    where cod_almacen = @cod_almacen 
      and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
    group by cod_producto;
  update sat
    set sat.existencia = sat.existencia + h.cantidad
    from Saldos_Productos_Almacen_Table sat, #cAcumulada h
    where sat.cod_producto = h.cod_producto;       

  update s
    set s.precio_costo_prom = round(p.pcosto / op.no_operac, 2)
    from Saldos_Productos_Almacen_Table s, no_operaciones op, costos_promedios p
    where s.cod_producto = op.cod_producto
      and s.cod_producto = p.cod_producto;      
/* FIN DEL CALCULO DEL PRECIO DE COSTO PROMEDIO */
      
  update Saldos_Productos_Almacen_Table set importe_costo_prom = existencia * precio_costo_prom;
  update Saldos_Productos_Almacen_Table set importe_venta = existencia * precio_venta;  
    
  select cod_producto, descripcion, u_medida, existencia as cant_acumulada, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta 
    from Saldos_Productos_Almacen_Table
    order by cod_producto;
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Rep_Costo_Mensual]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Costo_Mensual]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================

CREATE PROCEDURE [dbo].[Rep_Costo_Mensual] @NoConc varchar(4000)
AS
BEGIN
	
	SET NOCOUNT ON;
	
	/*
	create table #temp
	(cod_producto varchar(50),
	 descripcion varchar(256),
	 u_medida varchar(50),
	 cantidad money,
	 p_venta money,
	 p_costo money, 
	 imp_venta money,
	 imp_costo money,
	 gastosDDU money,
	 margen money,
	 porciento money
	)
	*/
	
	/*insert into #temp (cod_producto, descripcion, u_medida, p_venta, p_costo, cantidad, imp_venta, imp_costo,gastosDDU)*/
    select d.cod_producto, p.descripcion, p.u_medida, p.p_venta, cast(0.0 as money) as p_costo, SUM(d.cantidad) as cantidad, SUM(d.imp_venta) as imp_venta, cast(0.0 as money) as imp_costo,
             SUM(round(imp_venta * 0.003, 2)) as gastosDDU, cast(0.0 as money) as margen, cast(0.0 as money) as porciento
      into #temp             
      from TransDetalles d inner join Nom_Productos p on d.cod_producto = p.codigo 
      where d.imp_venta <> 0
        and CHARINDEX(d.documento + '','', @NoConc) <> 0
        and d.documento <> ''''
        and d.documento is not null
      group by d.cod_producto, p.descripcion, p.u_medida, p.p_venta, p.p_costo;
                            
/* CALCULO DEL PRECIO DE COSTO PROMEDIO */    
  declare @cod_almacen as varchar(10);          
      set @NoConc = SUBSTRING(@NoConc, 1, CHARINDEX('','', @NoConc) - 1);
      select distinct @cod_almacen = cod_almacen
        from TransDetalles 
        where documento = @NoConc
  truncate table no_operaciones;    
  insert into no_operaciones (cod_producto, no_operac)    
    select cod_producto, COUNT(*) as no_operac
      from TransDetalles 
      where cod_almacen = @cod_almacen 
        and (cod_transaccion = ''11'' or cod_transaccion = ''22'')
        and cod_producto in (select cod_producto from #temp)
      group by cod_producto;        
  select cod_producto, COUNT(*) as no_operac
    into #noOperac 
    from Entradas_Historico 
    where cod_almacen = @cod_almacen      
      and cod_producto in (select cod_producto from #temp)
    group by cod_producto;
  insert into No_Operaciones (cod_producto, no_operac)
    select cod_producto, 0
      from #noOperac
      where cod_producto not in (select cod_producto from No_Operaciones)  
  update no
    set no.no_operac = no.no_operac + h.no_operac
    from No_Operaciones no, #noOperac h
    where no.cod_producto = h.cod_producto;     
            
  truncate table costos_promedios;
  insert into costos_promedios (cod_producto, pcosto)    
    select cod_producto, SUM(round(pcosto,2)) as pcosto
      from TransDetalles 
      where cod_almacen = @cod_almacen 
        and (cod_transaccion = ''11'' or cod_transaccion = ''22'')
        and cod_producto in (select cod_producto from #temp)
      group by cod_producto;
  select cod_producto, SUM(p_costo) as pcosto
    into #pCosto
    from Entradas_Historico 
    where cod_almacen = @cod_almacen    
      and cod_producto in (select cod_producto from #temp)
    group by cod_producto; 
  insert into Costos_Promedios (cod_producto, pcosto)
    select cod_producto, 0
      from #pCosto  
      where cod_producto not in (select cod_producto from Costos_Promedios)  
  update pc
    set pc.pcosto = pc.pcosto + h.pcosto
    from Costos_Promedios pc, #pCosto h
    where pc.cod_producto = h.cod_producto;  
      
  update s
      set s.p_costo = p.pcosto / op.no_operac
      from #temp s, no_operaciones op, costos_promedios p
      where s.cod_producto = op.cod_producto
        and s.cod_producto = p.cod_producto;        
    
    update #temp set imp_costo = cantidad * p_costo;
/* FIN DEL CALCULO DEL PRECIO DE COSTO PROMEDIO */
                      
    update #temp 
      set margen = imp_venta - imp_costo,
          porciento = ((imp_venta - imp_costo) * 100) / imp_venta;
    select cod_producto, descripcion, u_medida, cantidad, p_venta, p_costo, imp_venta, imp_costo, gastosDDU, margen, porciento 
      from #temp 
      order by cod_producto;
     	  
/*     	  
	insert into #temp (cod_producto, descripcion, u_medida, p_venta, p_costo, cantidad, imp_venta, imp_costo,gastosDDU)
      select d.cod_producto, p.descripcion, p.u_medida, p.p_venta, p.p_costo, SUM(d.cantidad), SUM(d.imp_venta), SUM(d.imp_costo),
             SUM(round(imp_venta * 0.003, 2))
      from TransDetalles d inner join Nom_Productos p on d.cod_producto = p.codigo 
      where d.imp_venta <> 0
        and CHARINDEX(d.documento + '','', @NoConc) <> 0
        and d.documento <> ''''
        and d.documento is not null
      group by d.cod_producto, p.descripcion, p.u_medida, p.p_venta, p.p_costo;
    --update #temp set imp_costo = cantidad * p_costo, imp_venta = cantidad * p_venta; 
    update #temp 
      set --gastosDDU = round(imp_venta * 0.003, 2),
          margen = imp_venta - imp_costo,
          porciento = ((imp_venta - imp_costo) * 100) / imp_venta;
    select cod_producto, descripcion, u_medida, cantidad, p_venta, p_costo, imp_venta, imp_costo, gastosDDU, margen, porciento 
      from #temp 
      order by cod_producto;
*/     	  
     	  
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[BackUp_DB]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[BackUp_DB]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[BackUp_DB] @path varchar(256) 
AS
BEGIN
  
  set @path = @path + ''\SIM.bak'';

  BACKUP DATABASE SIM
   TO DISK = @path
     WITH FORMAT;

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Ad_Hoc_Distributed_Queries]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Ad_Hoc_Distributed_Queries]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Ad_Hoc_Distributed_Queries]
AS
BEGIN
	SET NOCOUNT ON;

    exec sp_configure ''Ad Hoc Distributed Queries'', 1
    reconfigure
    
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Saldos_Productos_Almacen]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Saldos_Productos_Almacen]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Saldos_Productos_Almacen] @cod_almacen varchar(50)
AS
BEGIN

	SET NOCOUNT ON;
	
	truncate table Saldos_Productos_Almacen_Table;
	/* Saldos Iniciales */
    insert into Saldos_Productos_Almacen_Table (cod_producto, descripcion, u_medida, saldo_inicial, entradas, salidas, existencia, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta)
      select s.cod_producto, p.descripcion, p.u_medida, SUM(s.saldo), 0, 0, 0, 0, 0, p.p_venta, 0
        from saldos s inner join Nom_Productos p on s.cod_producto = p.codigo 
        where s.cod_almacen = @cod_almacen 
        group by s.cod_producto, p.descripcion, p.u_medida, p.p_venta;
    
    /* Entradas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''1'' and cod_almacen = @cod_almacen)
           or (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and cod_almacen_origen <> '''' and cod_almacen_origen is not null)
        group by cod_producto;    
    update s
      set s.entradas = ent.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas ent  
      where s.cod_producto = ent.cod_producto
  
    /* Salidas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and (cod_almacen_origen = '''' or cod_almacen_origen is null))
        group by cod_producto;    
    update s
      set s.salidas = sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
    /* Salidas Transferencias */  
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''2'' and cod_almacen_origen = @cod_almacen)
        group by cod_producto;    
    update s
      set s.salidas = s.salidas + sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
            
    update Saldos_Productos_Almacen_Table set saldo_inicial = 0 where saldo_inicial is null;
    update Saldos_Productos_Almacen_Table set entradas = 0 where entradas is null;
    update Saldos_Productos_Almacen_Table set salidas = 0 where salidas is null;
    update Saldos_Productos_Almacen_Table set existencia = saldo_inicial + entradas - salidas;  
    update Saldos_Productos_Almacen_Table set existencia = 0 where existencia is null;
    
    delete from Saldos_Productos_Almacen_Table where existencia = 0;  
    
    update Saldos_Productos_Almacen_Table set importe_venta = existencia * precio_venta;  
    
    
  truncate table no_operaciones;
  insert into no_operaciones (cod_producto, no_operac)    
    select cod_producto, COUNT(*) as no_operac
      from TransDetalles 
      where cod_almacen = @cod_almacen 
        and (cod_transaccion = ''11'' or cod_transaccion = ''22'')
        and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
      group by cod_producto;
  select cod_producto, COUNT(*) as no_operac
    into #noOperac 
    from Entradas_Historico 
    where cod_almacen = @cod_almacen      
      and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
    group by cod_producto;
  insert into No_Operaciones (cod_producto, no_operac)
    select cod_producto, 0
      from #noOperac
      where cod_producto not in (select cod_producto from No_Operaciones)  
  update no
    set no.no_operac = no.no_operac + h.no_operac
    from No_Operaciones no, #noOperac h
    where no.cod_producto = h.cod_producto;     
            
  truncate table costos_promedios;
  insert into costos_promedios (cod_producto, pcosto)    
    select cod_producto, SUM(round(pcosto,2)) as pcosto
      from TransDetalles 
      where cod_almacen = @cod_almacen 
        and (cod_transaccion = ''11'' or cod_transaccion = ''22'')
        and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
      group by cod_producto;
  select cod_producto, SUM(p_costo) as pcosto
    into #pCosto
    from Entradas_Historico 
    where cod_almacen = @cod_almacen    
      and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
    group by cod_producto; 
  insert into Costos_Promedios (cod_producto, pcosto)
    select cod_producto, 0
      from #pCosto  
      where cod_producto not in (select cod_producto from Costos_Promedios)  
  update pc
    set pc.pcosto = pc.pcosto + h.pcosto
    from Costos_Promedios pc, #pCosto h
    where pc.cod_producto = h.cod_producto;
    
    
/*    
    truncate table no_operaciones;
    insert into no_operaciones (cod_producto, no_operac)    
      select cod_producto, COUNT(*) as no_operac
        from TransDetalles 
        where cod_transaccion = ''11''
          and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
        group by cod_producto;
        
    /*
    select cod_producto, COUNT(*) as no_operac
      into #temp_operac
      from Entradas_Historico  
      where cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
      group by cod_producto;
    update op
      set op.no_operac = op.no_operac + e.no_operac
      from no_operaciones op, #temp_operac e
      where op.cod_producto = e.cod_producto;
    drop table #temp_operac;  
    */
      
    truncate table costos_promedios;
    insert into costos_promedios (cod_producto, pcosto)    
      select cod_producto, SUM(pcosto) as pcosto
        from TransDetalles 
        where cod_transaccion = ''11''
          and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
        group by cod_producto;
        
    /*    
    select cod_producto, SUM(p_costo) as pcosto
      into #temp_costos
      from Entradas_Historico
      where cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
      group by cod_producto;
    update cp
      set cp.pcosto = cp.pcosto + c.pcosto
      from costos_promedios cp, #temp_costos c
      where cp.cod_producto = c.cod_producto;
    */
    */
    update s
      set s.precio_costo_prom = p.pcosto / op.no_operac
      from Saldos_Productos_Almacen_Table s, no_operaciones op, costos_promedios p
      where s.cod_producto = op.cod_producto
        and s.cod_producto = p.cod_producto;


    update Saldos_Productos_Almacen_Table set importe_costo_prom = existencia * precio_costo_prom;

    select cod_producto, descripcion, u_medida, existencia, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta 
      from Saldos_Productos_Almacen_Table
      order by cod_producto;
      
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Saldo_Un_Producto]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Saldo_Un_Producto]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Saldo_Un_Producto] @codigo varchar(50)
AS
BEGIN

	SET NOCOUNT ON;
	
  truncate table Temp_Saldo_Un_Producto;
  truncate table Entradas_Un_Producto;
  truncate table Salidas_Un_Producto;
  truncate table Salidas_Un_Producto_Transf;

  insert into Temp_Saldo_Un_Producto (cod_almacen, saldo_inicial, entradas, salidas, saldo_actual)
    select codigo, 0, 0, 0, 0
      from Nom_Almacenes ;
      
  /* Saldos Iniciales */
  update t
    set t.saldo_inicial = s.saldo 
    from Temp_Saldo_Un_Producto t, Saldos s    
    where t.cod_almacen = s.cod_almacen
      and s.cod_producto = @codigo;

  /* Entradas */    
  insert into Entradas_Un_Producto (cod_almacen, cantidad)  
    select cod_almacen, SUM(cantidad) as cantidad
      from TransDetalles     
      where cod_producto = @codigo 
        and ((left(cod_transaccion, 1) = ''1'')
          or (left(cod_transaccion, 1) = ''2'' and cod_almacen_origen <> '''' and cod_almacen_origen is not null))
      group by cod_almacen       
  update t
    set t.entradas = e.cantidad
    from Temp_Saldo_Un_Producto t, Entradas_Un_Producto e  
    where t.cod_almacen = e.cod_almacen;
    
  /* Salidas */  
  insert into Salidas_Un_Producto (cod_almacen, cantidad)  
    select cod_almacen, SUM(cantidad) as cantidad
      from TransDetalles     
      where cod_producto = @codigo 
        and (left(cod_transaccion, 1) = ''2'' and (cod_almacen_origen = '''' or cod_almacen_origen is null))
      group by cod_almacen           
  insert into Salidas_Un_Producto_Transf (cod_almacen, cantidad)  
    select cod_almacen_origen, SUM(cantidad) as cantidad
      from TransDetalles     
      where cod_producto = @codigo 
        and (left(cod_transaccion, 1) = ''2'' and (cod_almacen_origen <> '''' and cod_almacen_origen is not null))
      group by cod_almacen_origen                        
  update t
    set t.salidas = s.cantidad
    from Temp_Saldo_Un_Producto t, Salidas_Un_Producto s
    where t.cod_almacen = s.cod_almacen;
  update t
    set t.salidas = t.salidas + s.cantidad
    from Temp_Saldo_Un_Producto t, Salidas_Un_Producto_Transf s
    where t.cod_almacen = s.cod_almacen;
            
  update Temp_Saldo_Un_Producto set saldo_inicial = 0 where saldo_inicial is null;
  update Temp_Saldo_Un_Producto set entradas = 0 where entradas is null;
  update Temp_Saldo_Un_Producto set salidas = 0 where salidas is null;
  update Temp_Saldo_Un_Producto set saldo_actual = 0 where saldo_actual is null;
  update Temp_Saldo_Un_Producto set saldo_actual = saldo_inicial + entradas - salidas;  
    
  select cod_almacen, saldo_inicial, entradas, salidas, saldo_actual 
    from Temp_Saldo_Un_Producto;
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Rep_Rotacion_Mercancias]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Rotacion_Mercancias]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Rep_Rotacion_Mercancias] @cod_almacen varchar(50), @clasificaciones varchar(4000) 
AS
BEGIN

	SET NOCOUNT ON;

	truncate table Temp_Rotacion_Mercancias;
	
    declare @clasif varchar(6), @index integer, @pos integer, @fecha_inicial datetime, @fecha_final datetime;
    
    set @fecha_inicial = ''1/1/'' + CAST(YEAR(GETDATE()) AS VARCHAR);
    set @fecha_final = GETDATE();

    set @clasif = '''';
    set @index = 1;
    while charindex('','', @clasificaciones, @index) <> 0
      begin
        set @pos = charindex('','', @clasificaciones, @index)
        set @clasif = substring(@clasificaciones, @index, @pos - @index)  

  	    insert into Temp_Rotacion_Mercancias (cod_producto, descripcion, um, clasif, entradas, entradas_costo, entradas_venta, entradas_peso, 
  	                                                                                 salidas, salidas_costo, salidas_venta, salidas_peso,
  	                                                                                 saldo, saldo_costo, saldo_venta, saldo_peso, rotacion,
  	                                                                                 pcosto, pventa, peso)
	      select codigo, descripcion, u_medida, clasif, 0, 0, 0, 0, 
	                                                    0, 0, 0, 0,
	                                                    0, 0, 0, 0, 0,
	                                                    p_costo, p_venta, peso_especifico 
   	        from Nom_Productos 
	        where clasif = @clasif
	           or LEFT(clasif, 4) + ''00'' = @clasif 
	           or LEFT(clasif, 2) + ''0000'' = @clasif;
	        
        set @index = @pos + 1
	  end      
    
	truncate table Saldos_Productos_Almacen_Table;
	
	/* Saldos Iniciales */
    insert into Saldos_Productos_Almacen_Table (cod_producto, descripcion, u_medida, saldo_inicial, entradas, salidas, existencia, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta)
      select s.cod_producto, p.descripcion, p.u_medida, SUM(s.saldo), 0, 0, 0, 0, 0, p.p_venta, 0
        from saldos s inner join Nom_Productos p on s.cod_producto = p.codigo 
        where s.cod_almacen = @cod_almacen 
        group by s.cod_producto, p.descripcion, p.u_medida, p.p_venta;
    
    /* Entradas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''1'' and cod_almacen = @cod_almacen)
           or (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and cod_almacen_origen <> '''' and cod_almacen_origen is not null)
        group by cod_producto;    
    update s
      set s.entradas = ent.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas ent  
      where s.cod_producto = ent.cod_producto
  
    /* Salidas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and (cod_almacen_origen = '''' or cod_almacen_origen is null))
        group by cod_producto;    
    update s
      set s.salidas = sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
    /* Salidas Transferencias */  
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''2'' and cod_almacen_origen = @cod_almacen)
        group by cod_producto;    
    update s
      set s.salidas = s.salidas + sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
            
    update Saldos_Productos_Almacen_Table set saldo_inicial = 0 where saldo_inicial is null;
    update Saldos_Productos_Almacen_Table set entradas = 0 where entradas is null;
    update Saldos_Productos_Almacen_Table set salidas = 0 where salidas is null;
    update Saldos_Productos_Almacen_Table set existencia = saldo_inicial + entradas - salidas;  
    update Saldos_Productos_Almacen_Table set existencia = 0 where existencia is null;
	  
	update t
	  set t.saldo = s.existencia,
	      t.saldo_costo = s.existencia * t.pcosto,
	      t.saldo_venta = s.existencia * t.pventa,
	      t.saldo_peso = s.existencia * t.peso 
	  from Temp_Rotacion_Mercancias t, Saldos_Productos_Almacen_Table s
	  where t.cod_producto = s.cod_producto;
	  
	truncate table Entradas_Productos;
	
	insert into Entradas_Productos (cod_almacen, cod_producto, cantidad)
	  select @cod_almacen, cod_producto, SUM(cantidad)
	    from TransDetalles 
	    where cod_almacen = @cod_almacen 
	      and fecha >= @fecha_inicial
	      and fecha <= @fecha_final 
	      and cod_producto in (select cod_producto from Temp_Rotacion_Mercancias)
	      and (cod_transaccion = ''11'' or cod_transaccion = ''22'')
	    group by cod_producto;  
     update t     	      
       set t.entradas = e.cantidad,
           t.entradas_costo = e.cantidad * t.pcosto,
           t.entradas_venta = e.cantidad * t.pventa,
           t.entradas_peso = e.cantidad * t.peso 
       from Temp_Rotacion_Mercancias t, Entradas_Productos e
       where t.cod_producto = e.cod_producto; 

	truncate table Salidas_Productos;
	
	insert into Salidas_Productos (cod_almacen, cod_producto, cantidad)
	  select @cod_almacen, cod_producto, SUM(cantidad)
	    from TransDetalles 
	    where fecha >= @fecha_inicial
	      and fecha <= @fecha_final 
	      and cod_producto in (select cod_producto from Temp_Rotacion_Mercancias)
	      and ((cod_transaccion = ''21'' and cod_almacen = @cod_almacen) or (cod_transaccion = ''22'' and cod_almacen_origen = @cod_almacen))
        group by cod_producto;	      
     update t     	      
       set t.salidas = s.cantidad,
           t.salidas_costo = s.cantidad * t.pcosto,
           t.salidas_venta = s.cantidad * t.pventa,
           t.salidas_peso = s.cantidad * t.peso 
       from Temp_Rotacion_Mercancias t, Salidas_Productos s
       where t.cod_producto = s.cod_producto; 

/* <Costo de la Mercancia Vendidad> */
       
	truncate table Salidas_Productos;
	
	insert into Salidas_Productos (cod_almacen, cod_producto, cantidad)
	  select @cod_almacen, cod_producto, SUM(imp_costo)
	    from TransDetalles 
	    where fecha >= @fecha_inicial
	      and fecha <= @fecha_final 
	      and cod_producto in (select cod_producto from Temp_Rotacion_Mercancias)
	      and cod_transaccion = ''21'' 
	      and cod_almacen = @cod_almacen
        group by cod_producto;	      
     update t     	      
       set t.costo_dela_venta = s.cantidad
       from Temp_Rotacion_Mercancias t, Salidas_Productos s
       where t.cod_producto = s.cod_producto; 

/* </Costo de la Mercancia Vendidad> */

    declare @cod_producto varchar(50), @saldo_inicial money, @entradas_recep money, @entradas_transf money, @salidas_fact money, @salidas_transf money, @saldo money;
    declare t_cur cursor
      for select cod_producto
            from Temp_Rotacion_Mercancias
    open t_cur;
    fetch next from t_cur into @cod_producto;            
    while @@FETCH_STATUS = 0
    begin
      select @saldo_inicial = saldo from Saldos where cod_almacen = @cod_almacen and cod_producto = @cod_producto;
      if (@saldo_inicial is null) set @saldo_inicial = 0;
      select @entradas_recep = SUM(cantidad) 
        from TransDetalles 
          where cod_almacen = @cod_almacen 
            and cod_producto = @cod_producto
            and fecha < @fecha_inicial 
            and cod_transaccion = ''11'';
      if (@entradas_recep is null) set @entradas_recep = 0;    
      select @entradas_transf = SUM(cantidad) 
        from TransDetalles 
        where cod_almacen = @cod_almacen 
          and cod_producto = @cod_producto
          and fecha < @fecha_inicial 
          and cod_transaccion = ''22'';
      if (@entradas_transf is null) set @entradas_transf = 0;    
      select @salidas_fact = SUM(cantidad) 
        from TransDetalles 
        where cod_almacen = @cod_almacen 
          and cod_producto = @cod_producto
          and fecha < @fecha_inicial 
          and cod_transaccion = ''21'';
      if (@salidas_fact is null) set @salidas_fact = 0;    
      select @salidas_transf = SUM(cantidad) 
        from TransDetalles 
        where cod_almacen_origen = @cod_almacen 
          and cod_producto = @cod_producto
          and fecha < @fecha_inicial 
          and cod_transaccion = ''22'';
      if (@salidas_transf is null) set @salidas_transf = 0;    
      set @saldo = @saldo_inicial + (@entradas_recep + @entradas_transf) - (@salidas_fact + @salidas_transf); 
      update Temp_Rotacion_Mercancias set saldo_inicial = @saldo where cod_producto = @cod_producto;
      
      select @saldo_inicial = saldo from Saldos where cod_almacen = @cod_almacen and cod_producto = @cod_producto;
      if (@saldo_inicial is null) set @saldo_inicial = 0;
      select @entradas_recep = SUM(cantidad) 
        from TransDetalles 
          where cod_almacen = @cod_almacen 
            and cod_producto = @cod_producto
            and fecha < @fecha_final
            and cod_transaccion = ''11'';
      if (@entradas_recep is null) set @entradas_recep = 0;    
      select @entradas_transf = SUM(cantidad) 
        from TransDetalles 
        where cod_almacen = @cod_almacen 
          and cod_producto = @cod_producto
          and fecha < @fecha_final
          and cod_transaccion = ''22'';
      if (@entradas_transf is null) set @entradas_transf = 0;    
      select @salidas_fact = SUM(cantidad) 
        from TransDetalles 
        where cod_almacen = @cod_almacen 
          and cod_producto = @cod_producto
          and fecha < @fecha_final 
          and cod_transaccion = ''21'';
      if (@salidas_fact is null) set @salidas_fact = 0;    
      select @salidas_transf = SUM(cantidad) 
        from TransDetalles 
        where cod_almacen_origen = @cod_almacen 
          and cod_producto = @cod_producto
          and fecha < @fecha_final 
          and cod_transaccion = ''22'';
      if (@salidas_transf is null) set @salidas_transf = 0;    
      set @saldo = @saldo_inicial + (@entradas_recep + @entradas_transf) - (@salidas_fact + @salidas_transf); 
      update Temp_Rotacion_Mercancias set saldo_final = @saldo  where cod_producto = @cod_producto;

      fetch next from t_cur into @cod_producto;            
    end
    close t_cur;
    deallocate t_cur;

    update Temp_Rotacion_Mercancias
      set rotacion = (costo_dela_venta * 12 / MONTH(GETDATE())) / ((saldo_inicial + saldo_final) / 2)
      where (saldo_inicial + saldo_final) <> 0;
      
    update Temp_Rotacion_Mercancias set rotacion = 0 where rotacion is null;

	insert into Temp_Rotacion_Mercancias (descripcion, clasif, entradas_venta, entradas_peso, salidas_venta, salidas_peso, saldo_venta, saldo_peso)
	  select ''S U B T O T A L'', clasif + ''991'', SUM(ROUND(entradas_venta, 2)), SUM(ROUND(entradas_peso, 2)), 
	                                            SUM(ROUND(salidas_venta, 2)), SUM(ROUND(salidas_peso, 2)), 
	                                            SUM(ROUND(saldo_venta, 2)), SUM(ROUND(saldo_peso, 2))  
	    from Temp_Rotacion_Mercancias
	    group by clasif;
	insert into Temp_Rotacion_Mercancias (clasif)
	  select clasif + ''992''
	    from Temp_Rotacion_Mercancias
	    where LEN(clasif) = 6
	    group by clasif;
            
     delete from Temp_Rotacion_Mercancias where rotacion = 0 or rotacion is null;           
            
     select cod_producto, descripcion, um, entradas, entradas_costo, entradas_venta, entradas_peso,
                                           salidas, salidas_costo, salidas_venta, salidas_peso,
                                           saldo, saldo_costo, saldo_venta, saldo_peso, rotacion
       from Temp_Rotacion_Mercancias
       order by clasif, cod_producto;   

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Rep_Pedido_A_Fabrica]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Pedido_A_Fabrica]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Rep_Pedido_A_Fabrica] @cod_almacen varchar(50), @clasificaciones varchar(4000)
AS
BEGIN

	SET NOCOUNT ON;
	
	truncate table Temp_Propuesta_Fabrica;
	
    declare @clasif varchar(6), @index integer, @pos integer;

    set @clasif = '''';
    set @index = 1;
    while charindex('','', @clasificaciones, @index) <> 0
      begin
        set @pos = charindex('','', @clasificaciones, @index)
        set @clasif = substring(@clasificaciones, @index, @pos - @index)  

  	    insert into Temp_Propuesta_Fabrica (cod_producto, descripcion, um, clasif, saldo, minimo, propuesta, pcosto, imp_costo)
	      select codigo, descripcion, u_medida, clasif, 0, cant_minima, 0, p_costo, 0
   	        from Nom_Productos 
	        where clasif = @clasif
	           or LEFT(clasif, 4) + ''00'' = @clasif 
	           or LEFT(clasif, 2) + ''0000'' = @clasif;
	        
        set @index = @pos + 1
	  end      
	
	truncate table Saldos_Productos_Almacen_Table;
	
	/* Saldos Iniciales */
    insert into Saldos_Productos_Almacen_Table (cod_producto, descripcion, u_medida, saldo_inicial, entradas, salidas, existencia, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta)
      select s.cod_producto, p.descripcion, p.u_medida, SUM(s.saldo), 0, 0, 0, 0, 0, p.p_venta, 0
        from saldos s inner join Nom_Productos p on s.cod_producto = p.codigo 
        where s.cod_almacen = @cod_almacen 
        group by s.cod_producto, p.descripcion, p.u_medida, p.p_venta;
    
    /* Entradas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''1'' and cod_almacen = @cod_almacen)
           or (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and cod_almacen_origen <> '''' and cod_almacen_origen is not null)
        group by cod_producto;    
    update s
      set s.entradas = ent.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas ent  
      where s.cod_producto = ent.cod_producto
  
    /* Salidas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and (cod_almacen_origen = '''' or cod_almacen_origen is null))
        group by cod_producto;    
    update s
      set s.salidas = sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;

    /* Salidas Transferencias */  
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''2'' and cod_almacen_origen = @cod_almacen)
        group by cod_producto;    
    update s
      set s.salidas = s.salidas + sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
            
    update Saldos_Productos_Almacen_Table set saldo_inicial = 0 where saldo_inicial is null;
    update Saldos_Productos_Almacen_Table set entradas = 0 where entradas is null;
    update Saldos_Productos_Almacen_Table set salidas = 0 where salidas is null;
    update Saldos_Productos_Almacen_Table set existencia = saldo_inicial + entradas - salidas;  
    update Saldos_Productos_Almacen_Table set existencia = 0 where existencia is null;
	  
	update t
	  set t.saldo = s.existencia
	  from Temp_Propuesta_Fabrica t, Saldos_Productos_Almacen_Table s
	  where t.cod_producto = s.cod_producto;
	  
	update Temp_Propuesta_Fabrica set propuesta = minimo - saldo; 
	delete from Temp_Propuesta_Fabrica where propuesta <= 0;
	update Temp_Propuesta_Fabrica set imp_costo = propuesta * pcosto where propuesta is not null;

	insert into Temp_Propuesta_Fabrica (descripcion, clasif, imp_costo)
	  select ''S U B T O T A L'', clasif + ''991'', SUM(ROUND(imp_costo, 2))  
	    from Temp_Propuesta_Fabrica
	    group by clasif;
	insert into Temp_Propuesta_Fabrica (clasif)
	  select clasif + ''992''
	    from Temp_Propuesta_Fabrica
	    where LEN(clasif) = 6
	    group by clasif;
	  
    select cod_producto, descripcion, um, saldo, propuesta, pcosto, imp_costo
	  from Temp_Propuesta_Fabrica
	  where propuesta is not null and propuesta <> 0
	  order by clasif, cod_producto;
	
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Rep_Movimiento_Mercancias]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Movimiento_Mercancias]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'
-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Rep_Movimiento_Mercancias] @cod_almacen varchar(50), @clasificaciones varchar(4000), @fecha_inicial datetime, @fecha_final datetime 
AS
BEGIN

	SET NOCOUNT ON;

	truncate table Temp_Rotacion_Mercancias;
	
    declare @clasif varchar(6), @index integer, @pos integer;
    
    set @clasif = '''';
    set @index = 1;
    while charindex('','', @clasificaciones, @index) <> 0
      begin
        set @pos = charindex('','', @clasificaciones, @index)
        set @clasif = substring(@clasificaciones, @index, @pos - @index)  

  	    insert into Temp_Rotacion_Mercancias (cod_producto, descripcion, um, clasif, entradas, entradas_costo, entradas_venta, entradas_peso, 
  	                                                                                 salidas, salidas_costo, salidas_venta, salidas_peso,
  	                                                                                 saldo, saldo_costo, saldo_venta, saldo_peso, rotacion,
  	                                                                                 pcosto, pventa, peso)
	      select codigo, descripcion, u_medida, clasif, 0, 0, 0, 0, 
	                                                    0, 0, 0, 0,
	                                                    0, 0, 0, 0, 0,
	                                                    p_costo, p_venta, peso_especifico 
   	        from Nom_Productos 
	        where clasif = @clasif
	           or LEFT(clasif, 4) + ''00'' = @clasif 
	           or LEFT(clasif, 2) + ''0000'' = @clasif;
	        
        set @index = @pos + 1
	  end      
    
	truncate table Saldos_Productos_Almacen_Table;
	
	/* Saldos Iniciales */
    insert into Saldos_Productos_Almacen_Table (cod_producto, descripcion, u_medida, saldo_inicial, entradas, salidas, existencia, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta)
      select s.cod_producto, p.descripcion, p.u_medida, SUM(s.saldo), 0, 0, 0, 0, 0, p.p_venta, 0
        from saldos s inner join Nom_Productos p on s.cod_producto = p.codigo 
        where s.cod_almacen = @cod_almacen 
        group by s.cod_producto, p.descripcion, p.u_medida, p.p_venta;
    
    /* Entradas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where fecha <= @fecha_final 
          and ((left(cod_transaccion, 1) = ''1'' and cod_almacen = @cod_almacen)
              or (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and cod_almacen_origen <> '''' and cod_almacen_origen is not null))
        group by cod_producto;    
    update s
      set s.entradas = ent.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas ent  
      where s.cod_producto = ent.cod_producto
  
    /* Salidas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where fecha <= @fecha_final 
          and (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and (cod_almacen_origen = '''' or cod_almacen_origen is null))
        group by cod_producto;    
    update s
      set s.salidas = sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
    /* Salidas Transferencias */  
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where fecha <= @fecha_final 
          and (left(cod_transaccion, 1) = ''2'' and cod_almacen_origen = @cod_almacen)
        group by cod_producto;    
    update s
      set s.salidas = s.salidas + sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
            
    update Saldos_Productos_Almacen_Table set saldo_inicial = 0 where saldo_inicial is null;
    update Saldos_Productos_Almacen_Table set entradas = 0 where entradas is null;
    update Saldos_Productos_Almacen_Table set salidas = 0 where salidas is null;
    update Saldos_Productos_Almacen_Table set existencia = saldo_inicial + entradas - salidas;  
    update Saldos_Productos_Almacen_Table set existencia = 0 where existencia is null;
	  
	update t
	  set t.saldo = s.existencia,
	      t.saldo_costo = s.existencia * t.pcosto,
	      t.saldo_venta = s.existencia * t.pventa,
	      t.saldo_peso = s.existencia * t.peso 
	  from Temp_Rotacion_Mercancias t, Saldos_Productos_Almacen_Table s
	  where t.cod_producto = s.cod_producto;
	  
	truncate table Entradas_Productos;
	
	insert into Entradas_Productos (cod_almacen, cod_producto, cantidad)
	  select @cod_almacen, cod_producto, SUM(cantidad)
	    from TransDetalles 
	    where cod_almacen = @cod_almacen 
	      and fecha >= @fecha_inicial
	      and fecha <= @fecha_final 
	      and cod_producto in (select cod_producto from Temp_Rotacion_Mercancias)
	      and (cod_transaccion = ''11'' or cod_transaccion = ''22'')
	    group by cod_producto;  
     update t     	      
       set t.entradas = e.cantidad,
           t.entradas_costo = e.cantidad * t.pcosto,
           t.entradas_venta = e.cantidad * t.pventa,
           t.entradas_peso = e.cantidad * t.peso 
       from Temp_Rotacion_Mercancias t, Entradas_Productos e
       where t.cod_producto = e.cod_producto; 

	truncate table Salidas_Productos;
	
	insert into Salidas_Productos (cod_almacen, cod_producto, cantidad)
	  select @cod_almacen, cod_producto, SUM(cantidad)
	    from TransDetalles 
	    where fecha >= @fecha_inicial
	      and fecha <= @fecha_final 
	      and cod_producto in (select cod_producto from Temp_Rotacion_Mercancias)
	      and ((cod_transaccion = ''21'' and cod_almacen = @cod_almacen) or (cod_transaccion = ''22'' and cod_almacen_origen = @cod_almacen))
        group by cod_producto;	      
     update t     	      
       set t.salidas = s.cantidad,
           t.salidas_costo = s.cantidad * t.pcosto,
           t.salidas_venta = s.cantidad * t.pventa,
           t.salidas_peso = s.cantidad * t.peso 
       from Temp_Rotacion_Mercancias t, Salidas_Productos s
       where t.cod_producto = s.cod_producto; 

    delete from Temp_Rotacion_Mercancias 
      where entradas = 0 and entradas_venta = 0
        and salidas = 0 and salidas_venta = 0
        and saldo = 0 and saldo_venta = 0;

	insert into Temp_Rotacion_Mercancias (descripcion, clasif, entradas_venta, entradas_peso, salidas_venta, salidas_peso, saldo_venta, saldo_peso)
	  select ''S U B T O T A L'', clasif + ''991'', SUM(ROUND(entradas_venta, 2)), SUM(ROUND(entradas_peso, 2)), 
	                                            SUM(ROUND(salidas_venta, 2)), SUM(ROUND(salidas_peso, 2)), 
	                                            SUM(ROUND(saldo_venta, 2)), SUM(ROUND(saldo_peso, 2))  
	    from Temp_Rotacion_Mercancias
	    group by clasif;
	insert into Temp_Rotacion_Mercancias (clasif)
	  select clasif + ''992''
	    from Temp_Rotacion_Mercancias
	    where LEN(clasif) = 6
	    group by clasif;
            
     /* Esto es un parche para sumarle a las entradas los saldos iniciales */
     update t
       set t.entradas = s.saldo + t.entradas,
           t.entradas_venta = t.entradas_venta + (s.saldo * t.pventa)
       from Temp_Rotacion_Mercancias t, Saldos s
       where t.cod_producto = s.cod_producto 
         and s.cod_almacen = @cod_almacen;
            
     select cod_producto, descripcion, um, entradas, entradas_venta, 
                                           salidas, salidas_venta, 
                                           saldo, saldo_venta
       from Temp_Rotacion_Mercancias
       order by clasif, cod_producto;   

END

' 
END
GO
/****** Object:  StoredProcedure [dbo].[Rep_Comprativo_Inventarios]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Comprativo_Inventarios]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Rep_Comprativo_Inventarios] @cod_almacen varchar(50)
AS
BEGIN

  SET NOCOUNT ON;

	truncate table Saldos_Productos_Almacen_Table;
	
	/* Saldos Iniciales */
    insert into Saldos_Productos_Almacen_Table (cod_producto, descripcion, u_medida, saldo_inicial, entradas, salidas, existencia, precio_costo_prom, importe_costo_prom, precio_venta, importe_venta)
      select s.cod_producto, p.descripcion, p.u_medida, SUM(s.saldo), 0, 0, 0, 0, 0, p.p_venta, 0
        from saldos s inner join Nom_Productos p on s.cod_producto = p.codigo 
        where s.cod_almacen = @cod_almacen 
        group by s.cod_producto, p.descripcion, p.u_medida, p.p_venta;
    
    /* Entradas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''1'' and cod_almacen = @cod_almacen)
           or (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and cod_almacen_origen <> '''' and cod_almacen_origen is not null)
        group by cod_producto;    
    update s
      set s.entradas = ent.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas ent  
      where s.cod_producto = ent.cod_producto
  
    /* Salidas */
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''2'' and cod_almacen = @cod_almacen and (cod_almacen_origen = '''' or cod_almacen_origen is null))
        group by cod_producto;    
    update s
      set s.salidas = sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
    /* Salidas Transferencias */  
    truncate table entradas_salidas;
    insert into entradas_salidas (cod_producto, cantidad)
      select cod_producto, SUM(cantidad) as cantidad
        from TransDetalles     
        where (left(cod_transaccion, 1) = ''2'' and cod_almacen_origen = @cod_almacen)
        group by cod_producto;    
    update s
      set s.salidas = s.salidas + sal.cantidad
      from Saldos_Productos_Almacen_Table s, entradas_salidas sal
      where s.cod_producto = sal.cod_producto;
            
    update Saldos_Productos_Almacen_Table set saldo_inicial = 0 where saldo_inicial is null;
    update Saldos_Productos_Almacen_Table set entradas = 0 where entradas is null;
    update Saldos_Productos_Almacen_Table set salidas = 0 where salidas is null;
    update Saldos_Productos_Almacen_Table set existencia = saldo_inicial + entradas - salidas;  
    update Saldos_Productos_Almacen_Table set existencia = 0 where existencia is null;
    
    delete from Saldos_Productos_Almacen_Table where existencia = 0;  
    
    update Saldos_Productos_Almacen_Table set importe_venta = existencia * precio_venta;  
    
    truncate table no_operaciones;
    insert into no_operaciones (cod_producto, no_operac)    
      select cod_producto, COUNT(*) as no_operac
        from TransDetalles 
        where cod_transaccion = ''11''
          and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
        group by cod_producto;
      
    truncate table costos_promedios;
    insert into costos_promedios (cod_producto, pcosto)    
      select cod_producto, SUM(pcosto) as pcosto
        from TransDetalles 
        where cod_transaccion = ''11''
          and cod_producto in (select cod_producto from Saldos_Productos_Almacen_Table)
        group by cod_producto;
    
    update s
      set s.precio_costo_prom = p.pcosto / op.no_operac
      from Saldos_Productos_Almacen_Table s, no_operaciones op, costos_promedios p
      where s.cod_producto = op.cod_producto
        and s.cod_producto = p.cod_producto;
      
    update Saldos_Productos_Almacen_Table set importe_costo_prom = existencia * precio_costo_prom;

    truncate table temp_comparacion_saldos;
  
    insert into temp_comparacion_saldos (cod_producto, descripcion, um, saldo_sistema, saldo_importado, diferencia)
      select s.cod_producto, p.descripcion, p.u_medida, 0, i.saldo, 0
        from Saldos s inner join Nom_Productos p on s.cod_producto = p.codigo 
                        inner join XLSImport i on s.cod_almacen = i.cod_almacen and p.codigo = i.cod_producto
        where s.cod_almacen = @cod_almacen;                              
             
    update t
      set t.saldo_sistema = s.existencia,
          t.diferencia = s.existencia - t.saldo_importado                           
      from temp_comparacion_saldos t, Saldos_Productos_Almacen_Table s
      where t.cod_producto = s.cod_producto;
                             
    select cod_producto, descripcion, um, saldo_sistema, saldo_importado, diferencia
      from temp_comparacion_saldos
      order by cod_producto;
    
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Rep_Movimientos_Producto]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Movimientos_Producto]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Rep_Movimientos_Producto] @cod_almacen varchar(50), @cod_producto varchar(50), @fecha_inicial datetime, @fecha_final datetime
AS
BEGIN
	SET NOCOUNT ON;

  truncate table Movimientos_Producto;

  declare @saldo_inicial money, @entradas_recep money, @entradas_transf money, @salidas_fact money, @salidas_transf money, @saldo money;
  
  select @saldo_inicial = saldo from Saldos where cod_almacen = @cod_almacen and cod_producto = @cod_producto;
  if (@saldo_inicial is null) set @saldo_inicial = 0;
  select @entradas_recep = SUM(cantidad) 
    from TransDetalles 
    where cod_almacen = @cod_almacen 
      and cod_producto = @cod_producto
      and fecha < @fecha_inicial 
      and cod_transaccion = ''11'';
  if (@entradas_recep is null) set @entradas_recep = 0;    
  select @entradas_transf = SUM(cantidad) 
    from TransDetalles 
    where cod_almacen = @cod_almacen 
      and cod_producto = @cod_producto
      and fecha < @fecha_inicial 
      and cod_transaccion = ''22'';
  if (@entradas_transf is null) set @entradas_transf = 0;    
  select @salidas_fact = SUM(cantidad) 
    from TransDetalles 
    where cod_almacen = @cod_almacen 
      and cod_producto = @cod_producto
      and fecha < @fecha_inicial 
      and cod_transaccion = ''21'';
  if (@salidas_fact is null) set @salidas_fact = 0;    
  select @salidas_transf = SUM(cantidad) 
    from TransDetalles 
    where cod_almacen_origen = @cod_almacen 
      and cod_producto = @cod_producto
      and fecha < @fecha_inicial 
      and cod_transaccion = ''22'';
  if (@salidas_transf is null) set @salidas_transf = 0;    
  set @saldo = @saldo_inicial + (@entradas_recep + @entradas_transf) - (@salidas_fact + @salidas_transf); 

  insert into Movimientos_Producto (Transaccion, Numero, Fecha, Referencia, Entrada, Salida, Saldo)
    select ''Recepción'', numero, fecha, documento, cantidad, 0, 0
      from TransDetalles
      where cod_almacen = @cod_almacen 
        and cod_producto = @cod_producto 
        and fecha >= @fecha_inicial
        and fecha <= @fecha_final
        and cod_transaccion = ''11'';
        
  insert into Movimientos_Producto (Transaccion, Numero, Fecha, Referencia, Entrada, Salida, Saldo)
    select ''Transferencia de Entrada'', numero, fecha, documento, cantidad, 0, 0
      from TransDetalles
      where cod_almacen = @cod_almacen 
        and cod_producto = @cod_producto 
        and fecha >= @fecha_inicial
        and fecha <= @fecha_final
        and cod_transaccion = ''22'';
      
  insert into Movimientos_Producto (Transaccion, Numero, Fecha, Referencia, Entrada, Salida, Saldo)
    select ''Transferencia de Salida'', numero, fecha, documento, 0, cantidad, 0
      from TransDetalles
      where cod_almacen_origen = @cod_almacen 
        and cod_producto = @cod_producto 
        and fecha >= @fecha_inicial
        and fecha <= @fecha_final
        and cod_transaccion = ''22'';

  insert into Movimientos_Producto (Transaccion, Numero, Fecha, Referencia, Entrada, Salida, Saldo)
    select ''Factura'', numero, fecha, documento, 0, cantidad, 0
      from TransDetalles
      where cod_almacen = @cod_almacen 
        and cod_producto = @cod_producto 
        and fecha >= @fecha_inicial
        and fecha <= @fecha_final
        and cod_transaccion = ''21'';

   declare @transaccion varchar(50), @numero int, @fecha datetime, @referencia varchar(50), @entrada money, @salida money;     
   declare t_cur cursor
      for select Transaccion, Numero, Fecha, Referencia, Entrada, Salida
            from Movimientos_Producto
            order by Fecha;
   open t_cur;
   fetch next from t_cur into @transaccion, @numero, @fecha, @referencia, @entrada, @salida;
   while @@FETCH_STATUS = 0
   begin
     update Movimientos_Producto
       set Saldo = @saldo + @entrada - @salida 
       where transaccion = @transaccion 
         and numero = @numero 
         and fecha = @fecha
         and referencia = @referencia;
     set @saldo = @saldo + @entrada - @salida;
     fetch next from t_cur into @transaccion, @numero, @fecha, @referencia, @entrada, @salida;
   end
   close t_cur;
   deallocate t_cur;         
                
   select Transaccion, Numero, Fecha, Referencia, Entrada, Salida, Saldo       
     from Movimientos_Producto
     order by fecha;
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Rep_Costo_Mensual_Indicadores]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Costo_Mensual_Indicadores]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================

CREATE PROCEDURE [dbo].[Rep_Costo_Mensual_Indicadores] @NoConc varchar(4000)
AS
BEGIN
	SET NOCOUNT ON;
	
	truncate table Temp_Costo_Mensual_Indicadores;
	
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 1, ''IMPORTE TOTAL DDU FACURADO EN EL MES'', SUM(imp_venta)
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
/*
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 2, ''IMPORTE TOTAL AL COSTO DDU DEL MES'', SUM(imp_costo) + SUM(round(imp_venta * 0.003, 2))
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null    
*/          
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 2, ''IMPORTE TOTAL AL COSTO DDU DEL MES'', SUM(d.cantidad * (c.pcosto / o.no_operac)) --+ SUM(round(d.imp_venta * 0.003, 2))
	    from TransDetalles d inner join Costos_Promedios c on d.cod_producto = c.cod_producto 
	                           inner join No_Operaciones o on d.cod_producto = o.cod_producto 
        where CHARINDEX(d.documento + '','', @NoConc) <> 0
          and d.documento <> ''''
          and d.documento is not null    
          
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 3, p.descripcion, SUM((imp_venta * p.valor) / 100)
	    from TransDetalles, Parametros p 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
          and p.codigo = ''C1''
        group by p.descripcion  
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 4, p.descripcion, SUM((imp_venta * p.valor) / 100)
	    from TransDetalles, Parametros p 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
          and p.codigo = ''C2''
        group by p.descripcion  
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 5, ''TOTAL DEL RESTO DE LOS GASTOS'', SUM(Importe)
	    from Facturas_Gasto  
        where CHARINDEX(NoConc + '','', @NoConc) <> 0
          and NoConc <> ''''
          and NoConc is not null
          
   declare @v1 money;       
   select @v1 = SUM((imp_venta * p.valor) / 100)
	    from TransDetalles, Parametros p 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
          and p.codigo = ''C1''
   declare @v2 money;       
   select @v2 = SUM((imp_venta * p.valor) / 100)
	    from TransDetalles, Parametros p 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
          and p.codigo = ''C2''
   declare @v3 money;       
   select @v3 = SUM(Importe)
	    from Facturas_Gasto  
        where CHARINDEX(NoConc + '','', @NoConc) <> 0
          and NoConc <> ''''
          and NoConc is not null
    if (@v1 is null) set @v1 = 0;            
    if (@v2 is null) set @v2 = 0;            
    if (@v3 is null) set @v3 = 0;            
    
/*        
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 2, ''IMPORTE TOTAL AL COSTO DDU DEL MES'', SUM(d.cantidad * (c.pcosto / o.no_operac)) + SUM(round(d.imp_venta * 0.003, 2))
	    from TransDetalles d inner join Costos_Promedios c on d.cod_producto = c.cod_producto 
	                           inner join No_Operaciones o on d.cod_producto = o.cod_producto 
        where CHARINDEX(d.documento + '','', @NoConc) <> 0
          and d.documento <> ''''
          and d.documento is not null    
*/
    
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 6, ''MARGEN BRUTO SOBRE VENTAS DDU'', SUM(d.imp_venta)
	                                             - SUM(d.cantidad * (c.pcosto / o.no_operac)) --+ SUM(round(d.imp_venta * 0.003, 2)) 
	                                             - @v1 
	                                             - @v2
	                                             - @v3
	    from TransDetalles d inner join Costos_Promedios c on d.cod_producto = c.cod_producto 
	                           inner join No_Operaciones o on d.cod_producto = o.cod_producto 
        where CHARINDEX(d.documento + '','', @NoConc) <> 0
          and d.documento <> ''''
          and d.documento is not null
          
/*    
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 6, ''MARGEN BRUTO SOBRE VENTAS DDU'', SUM(imp_venta) 
	                                               - (SUM(imp_costo) + SUM(round(imp_venta * 0.003, 2)))
	                                               - @v1 
	                                               - @v2
	                                               - @v3
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
*/

	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 7, ''MARGEN BRUTO SOBRE VENTAS DDU EN %'', ((SUM(d.imp_venta) 
	                                                      - SUM(d.cantidad * (c.pcosto / o.no_operac)) --+ SUM(round(d.imp_venta * 0.003, 2)) 
	                                                      - @v1 
	                                                      - @v2
	                                                      - @v3) * 100) / SUM(d.imp_venta)
	    from TransDetalles d inner join Costos_Promedios c on d.cod_producto = c.cod_producto 
	                           inner join No_Operaciones o on d.cod_producto = o.cod_producto 
        where CHARINDEX(d.documento + '','', @NoConc) <> 0
          and d.documento <> ''''
          and d.documento is not null

/*          
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 7, ''MARGEN BRUTO SOBRE VENTAS DDU EN %'', ((SUM(imp_venta) 
	                                                      - (SUM(imp_costo) + SUM(round(imp_venta * 0.003, 2)))
	                                                      - @v1 
	                                                      - @v2
	                                                      - @v3) * 100) / SUM(imp_venta)
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
*/

    update Temp_Costo_Mensual_Indicadores set importe = 0 where importe is null;

	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  values (8, ''INGRESO TOTAL CANPLASTICA DEL MES USD'', 0.0)
	update t
	  set t.importe = t1.importe - t2.importe - t3.importe - t4.importe
	  from Temp_Costo_Mensual_Indicadores t, 
	       Temp_Costo_Mensual_Indicadores t1,
	       Temp_Costo_Mensual_Indicadores t2,
	       Temp_Costo_Mensual_Indicadores t3,
	       Temp_Costo_Mensual_Indicadores t4
	  where t.id = 8
	    and t1.id = 1
	    and t2.id = 3
	    and t3.id = 4
	    and t4.id = 5

    update Temp_Costo_Mensual_Indicadores set importe = 0 where importe is null;
/*
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 9, ''GASTOS DDU DE ALMACENAJE (3 %)'', SUM(imp_venta) * 0.3 / 100
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null

	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  values (10, ''MARGEN DEL INGRESO SOBRE VENTAS DDU'', 0.0)
	update t
	  set t.importe = t1.importe - t2.importe - t3.importe
	  from Temp_Costo_Mensual_Indicadores t, 
	       Temp_Costo_Mensual_Indicadores t1,
	       Temp_Costo_Mensual_Indicadores t2,
	       Temp_Costo_Mensual_Indicadores t3
	  where t.id = 10
	    and t1.id = 8
	    and t2.id = 9
	    and t3.id = 2

	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  values (11, ''MARGEN DEL INGRESO SOBRE VENTAS DDU EN %'', 0.0)
	update t
	  set t.importe = (t1.importe * 100) / t2.importe
	  from Temp_Costo_Mensual_Indicadores t, 
	       Temp_Costo_Mensual_Indicadores t1,
	       Temp_Costo_Mensual_Indicadores t2
	  where t.id = 11
	    and t1.id = 10
	    and t2.id = 1
*/
    select descripcion, importe
      from Temp_Costo_Mensual_Indicadores
      order by id;
	
END


/*
ALTER PROCEDURE [dbo].[Rep_Costo_Mensual_Indicadores] @NoConc varchar(4000)
AS
BEGIN
	SET NOCOUNT ON;
	
	truncate table Temp_Costo_Mensual_Indicadores;
	
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 1, ''IMPORTE TOTAL DDU FACURADO EN EL MES'', SUM(imp_venta)
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 2, ''IMPORTE TOTAL AL COSTO CIF DEL MES'', SUM(imp_costo)
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 3, ''MARGEN BRUTO SOBRE VENTAS DDU'', SUM(imp_venta) - SUM(imp_costo)
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 4, ''MARGEN BRUTO SOBRE VENTAS DDU EN %'', ((SUM(imp_venta) - SUM(imp_costo)) * 100) / SUM(imp_venta)
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 5, ''COMISION 5% de AUSA'', SUM((imp_venta * 5) / 100)
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 6, ''COMISION 4% de AUSA'', SUM((imp_venta * 4) / 100)
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 7, ''TOTAL DEL RESTO DE LOS GASTOS'', SUM(Importe)
	    from Facturas_Gasto  
        where CHARINDEX(NoConc + '','', @NoConc) <> 0
          and NoConc <> ''''
          and NoConc is not null

    update Temp_Costo_Mensual_Indicadores set importe = 0 where importe is null;

	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  values (8, ''INGRESO TOTAL CANPLASTICA DEL MES USD'', 0.0)
	update t
	  set t.importe = t1.importe - t2.importe - t3.importe - t4.importe
	  from Temp_Costo_Mensual_Indicadores t, 
	       Temp_Costo_Mensual_Indicadores t1,
	       Temp_Costo_Mensual_Indicadores t2,
	       Temp_Costo_Mensual_Indicadores t3,
	       Temp_Costo_Mensual_Indicadores t4
	  where t.id = 8
	    and t1.id = 1
	    and t2.id = 5
	    and t3.id = 6
	    and t4.id = 7     

    update Temp_Costo_Mensual_Indicadores set importe = 0 where importe is null;

	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  select 9, ''GASTOS DDU DE ALMACENAJE (3 %)'', SUM(imp_venta) * 0.3 / 100
	    from TransDetalles 
        where CHARINDEX(documento + '','', @NoConc) <> 0
          and documento <> ''''
          and documento is not null

	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  values (10, ''MARGEN DEL INGRESO SOBRE VENTAS DDU'', 0.0)
	update t
	  set t.importe = t1.importe - t2.importe - t3.importe
	  from Temp_Costo_Mensual_Indicadores t, 
	       Temp_Costo_Mensual_Indicadores t1,
	       Temp_Costo_Mensual_Indicadores t2,
	       Temp_Costo_Mensual_Indicadores t3
	  where t.id = 10
	    and t1.id = 8
	    and t2.id = 9
	    and t3.id = 2
	insert into Temp_Costo_Mensual_Indicadores (id, descripcion, importe)
	  values (11, ''MARGEN DEL INGRESO SOBRE VENTAS DDU EN %'', 0.0)
	update t
	  set t.importe = (t1.importe * 100) / t2.importe
	  from Temp_Costo_Mensual_Indicadores t, 
	       Temp_Costo_Mensual_Indicadores t1,
	       Temp_Costo_Mensual_Indicadores t2
	  where t.id = 11
	    and t1.id = 10
	    and t2.id = 1
    select descripcion, importe
      from Temp_Costo_Mensual_Indicadores
      order by id;
	
END
*/' 
END
GO
/****** Object:  StoredProcedure [dbo].[Rep_Conciliaciones_Facturas]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Rep_Conciliaciones_Facturas]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Rep_Conciliaciones_Facturas] @almacenes varchar(256), @documento varchar(50)
AS
BEGIN
	SET NOCOUNT ON;

    select numero, fecha,  SUM(imp_venta) as [Importe (USD)], 
                           SUM((imp_venta * 5) / 100) as [Comisión 5%], 
                           SUM((imp_venta * 4) / 100) as [Comision 3%],
                           SUM((imp_venta - ((imp_venta * 5) / 100) - ((imp_venta * 4) / 100)) / 1) as [Importe a pagar (CUC)],
                           SUM(imp_venta - ((imp_venta * 5) / 100) - ((imp_venta * 4) / 100)) as [Importe a pagar (USD)]
       from TransDetalles 
       where cod_transaccion = ''21'' 
         and CHARINDEX(cod_almacen + '','', @almacenes) <> 0
         and documento = @documento 
       group by numero, fecha

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Recalcular_Recepcion]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Recalcular_Recepcion]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Recalcular_Recepcion] @cod_almacen varchar(50), @documento varchar(50), @numero int, @fecha datetime,
                                     @total_recepcion money, @total_impuestos money
AS
BEGIN

  update TransDetalles set incremento = 0 where incremento is null;
  declare @total money;
  select @total = SUM((pcosto - incremento) * cantidad)
    from TransDetalles 
    where cod_almacen = @cod_almacen
      and documento = @documento
      and numero = @numero
      and fecha = @fecha;

  declare @porciento money;
  if (@total <> 0 and @total is not null)
    set @porciento = (@total_impuestos * 100) / @total; 
    else set @porciento = 0;
  
  declare @id int, @cod_producto varchar(50), @pcosto money, @incrementoAct money, @incremento money;
  declare recep_cur cursor
    for select id, cod_producto, pcosto, incremento
          from TransDetalles
          where cod_almacen = @cod_almacen
            and documento = @documento
            and numero = @numero
            and fecha = @fecha;
  open recep_cur;
  fetch next from recep_cur into @id, @cod_producto, @pcosto, @incremento;
  while @@FETCH_STATUS = 0
  begin
    if (@incremento is null) set @incremento = 0;
    set @incrementoAct = round(((@pcosto - @incremento) * @porciento) / 100, 2);
    update TransDetalles 
      set pcosto = (@pcosto - @incremento) + @incrementoAct,
          imp_costo = ((@pcosto - @incremento) + @incrementoAct) * cantidad,
          incremento = @incrementoAct 
      where id = @id;    
    fetch next from recep_cur into @id, @cod_producto, @pcosto, @incremento;
  end
  close recep_cur;
  deallocate recep_cur;            

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[DeleteDetalles]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[DeleteDetalles]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[DeleteDetalles] @cod_almacen varchar(50), @documento varchar(50), @numero int, @fecha datetime, @cod_transaccion varchar(2)
AS
BEGIN

  delete from TransDetalles 
  where cod_almacen = @cod_almacen
    and documento = @documento
    and numero = @numero
    and fecha = @fecha
    and cod_transaccion = @cod_transaccion;

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Total_ImporteCUC_Facturas]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Total_ImporteCUC_Facturas]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Total_ImporteCUC_Facturas] @NoConc varchar(50)
AS
BEGIN
	SET NOCOUNT ON;
	
	truncate table Temp_TotalCUC_Facturas;
	truncate table Temp_TotalFact_Gasto;
	
	insert into Temp_TotalCUC_Facturas (ImporteCUC)
      select SUM((imp_venta - ((imp_venta * 5) / 100) - ((imp_venta * 4) / 100)))
        from TransDetalles
        where cod_transaccion = ''21'' 
          and documento = @NoConc;
	
	insert into Temp_TotalFact_Gasto (ImporteCUC)
  	  select SUM(Importe)
  	    from Facturas_Gasto 
  	    where NoConc = @NoConc;
	  
	if exists (select ImporteCUC from Temp_TotalFact_Gasto where ImporteCUC is not null)
	begin  
  	  update t
	     set t.ImporteCUC = t.ImporteCUC - f.ImporteCUC
	     from Temp_TotalCUC_Facturas t, Temp_TotalFact_Gasto f
	end
	else begin
    	   update t
	          set t.ImporteCUC = t.ImporteCUC
	          from Temp_TotalCUC_Facturas t
	     end     
	
    select ImporteCUC
      from Temp_TotalCUC_Facturas 
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[UpdateDetalles]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[UpdateDetalles]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[UpdateDetalles] @cod_almacen varchar(50), @documento varchar(50), @numero int, @fecha datetime, @cod_transaccion varchar(2),
                                @old_cod_almacen varchar(50), @old_documento varchar(50), @old_numero int, @old_fecha datetime, @old_cod_transaccion varchar(2)
AS
BEGIN

  update TransDetalles 
    set cod_almacen = @cod_almacen,
        documento = @documento,
        numero = @numero,
        fecha = @fecha,
        cod_transaccion = @cod_transaccion
    where cod_almacen = @old_cod_almacen
      and documento = @old_documento
      and numero = @old_numero
      and fecha = @old_fecha 
      and cod_transaccion = @old_cod_transaccion;    

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Utilitarios_BorrarTablas]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Utilitarios_BorrarTablas]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Utilitarios_BorrarTablas]
AS
BEGIN
  truncate table Facturas_Gasto
  truncate table Impuestos
  --truncate table Nom_Ajustes
  truncate table Nom_Almacenes
  truncate table Nom_Clasif_Producto
  truncate table Nom_Clientes
  truncate table Nom_Productos
  truncate table Nom_Proveedores
  --truncate table Nom_Transacciones
  truncate table Saldos
  truncate table TransDetalles
  truncate table TransEncabezados
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Act_Precios_Recepcion]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Act_Precios_Recepcion]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Act_Precios_Recepcion] @id int, @euro bit, @tasa_cambio money 
AS
BEGIN
	SET NOCOUNT ON;
	
	declare @idDet int, @cod_producto varchar(50), @clasif varchar(6), @peso_especifico money, @p_adicional money, @descuento money,
	        @cod_almacen varchar(50), @documento varchar(50), @numero int, @fecha datetime, @p_costo money, @precio_costo money;
	select @cod_almacen = cod_almacen, @documento = documento, @numero = numero, @fecha = fecha
	  from TransEncabezados 
	  where id = @id;
	
	declare det_cur cursor
	  for select d.id, d.cod_producto, p.clasif, p.peso_especifico, p.p_adicional, p.descuento, p.p_costo
	        from TransDetalles d inner join Nom_Productos p on d.cod_producto = p.codigo
	        where d.cod_almacen = @cod_almacen
	          and d.documento = @documento
	          and d.numero = @numero
	          and d.fecha = @fecha
	          and d.cod_transaccion = ''11'';
    open det_cur;
    fetch next from det_cur into @idDet, @cod_producto, @clasif, @peso_especifico, @p_adicional, @descuento, @p_costo;
    while @@FETCH_STATUS = 0
    begin
      if (LEFT(@clasif, 2) = ''01'' or LEFT(@clasif, 2) = ''03'')
      begin
        set @precio_costo = @peso_especifico * @p_adicional * @tasa_cambio;
      end
      else begin
             set @precio_costo = @p_adicional * @descuento * @tasa_cambio;
           end
      if (@euro = 1)
      begin     
        update TransDetalles 
          set pcosto = @precio_costo, imp_costo = cantidad * @precio_costo 
          where id = @idDet;
      end
      else begin
             update TransDetalles 
               set pcosto = @p_costo, imp_costo = cantidad * @p_costo  
               where id = @idDet;
           end
      fetch next from det_cur into @idDet, @cod_producto, @clasif, @peso_especifico, @p_adicional, @descuento, @p_costo;
    end
    close det_cur;
    deallocate det_cur;

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Saldo_Productos]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Saldo_Productos]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Saldo_Productos]
AS
BEGIN

	SET NOCOUNT ON;

  truncate table Temp_Saldo_Productos;
  truncate table Entradas_Productos;
  truncate table Salidas_Productos;

  insert into Temp_Saldo_Productos (cod_almacen, cod_producto, saldo_inicial)
    select cod_almacen, cod_producto, SUM(saldo) 
      from saldos
      group by cod_almacen, cod_producto;
    
  insert into Entradas_Productos (cod_almacen, cod_producto, cantidad)  
    select cod_almacen, cod_producto, SUM(cantidad) as cantidad
      from TransDetalles     
      where left(cod_transaccion, 1) = ''1''
      group by cod_almacen, cod_producto;
    
  update t
    set t.entradas = e.cantidad
    from Temp_Saldo_Productos t, Entradas_Productos e  
    where t.cod_almacen = e.cod_almacen
      and t.cod_producto = e.cod_producto
  
  insert into Salidas_Productos (cod_almacen, cod_producto, cantidad)  
    select cod_almacen, cod_producto, SUM(cantidad) as cantidad
      from TransDetalles     
      where left(cod_transaccion, 1) = ''2''
      group by cod_almacen, cod_producto;
    
  update t
    set t.salidas = s.cantidad
    from Temp_Saldo_Productos t, Salidas_Productos s
    where t.cod_almacen = s.cod_almacen
      and t.cod_producto = s.cod_producto;
    
  update Temp_Saldo_Productos set saldo_actual = saldo_inicial + entradas - salidas;  
  update Temp_Saldo_Productos set saldo_inicial = 0 where saldo_inicial is null;
  update Temp_Saldo_Productos set entradas = 0 where entradas is null;
  update Temp_Saldo_Productos set salidas = 0 where salidas is null;
  update Temp_Saldo_Productos set saldo_actual = 0 where saldo_actual is null;
    
  select cod_almacen, cod_producto, saldo_inicial, entradas, salidas, saldo_actual 
    from Temp_Saldo_Productos;
END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Reporte1]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Reporte1]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Reporte1]
AS
BEGIN
	SET NOCOUNT ON;

    select numero, fecha,  SUM(imp_venta) as [Importe (USD)], 
                           SUM((imp_venta * 5) / 100) as [Comisión 5%], 
                           SUM((imp_venta * 3) / 100) as [Comision 3%],
                           SUM((imp_venta - ((imp_venta * 5) / 100) - ((imp_venta * 3) / 100)) / 1.08) as [Importe a pagar (CUC)],
                           SUM(imp_venta - ((imp_venta * 5) / 100) - ((imp_venta * 3) / 100)) as [Importe a pagar (USD)]
       from TransDetalles 
       where cod_transaccion = ''21'' 
         and cod_almacen = ''01''
         and documento = ''1-2010''
       group by numero, fecha

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Utilitarios_Clasif_Produsctos_Textos]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Utilitarios_Clasif_Produsctos_Textos]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Utilitarios_Clasif_Produsctos_Textos] 
AS
BEGIN
	SET NOCOUNT ON;

    truncate table Utilitarios_Clasif_Produsctos_Textos_Table;
    
    insert into Utilitarios_Clasif_Produsctos_Textos_Table (codigo, descripcion)
      select codigo, LTRIM(descripcion)
        from Nom_Clasif_Producto;
        
    update c
      set c.descripcion = LTRIM(c1.descripcion) + '' - '' + LTRIM(c.descripcion)
      from Utilitarios_Clasif_Produsctos_Textos_Table c, Nom_Clasif_Producto c1
      where left(c.codigo, 4) + ''00'' = c1.codigo
        and c.codigo <> c1.codigo;
       
    update c
      set c.descripcion = LTRIM(c1.descripcion) + '' - '' + LTRIM(c.descripcion)
      from Utilitarios_Clasif_Produsctos_Textos_Table c, Nom_Clasif_Producto c1
      where left(c.codigo, 2) + ''0000'' = c1.codigo
        and c.codigo <> c1.codigo;

    select codigo, descripcion
      from Utilitarios_Clasif_Produsctos_Textos_Table
      order by codigo;

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Utilitarios_Reconstruir_Un_Encabezado]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Utilitarios_Reconstruir_Un_Encabezado]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Utilitarios_Reconstruir_Un_Encabezado] @id int
AS
BEGIN

  declare @cod_almacen varchar(50),
          @documento varchar(50),
          @numero int,
          @fecha datetime,
          @cod_transaccion varchar(2)

  select @cod_almacen = cod_almacen, @documento = documento, @numero = numero, 
         @fecha = fecha, @cod_transaccion = cod_transaccion 
    from TransEncabezados
    where id = @id
    
  select SUM(cantidad) as cantidad, SUM(imp_costo) as imp_costo, SUM(imp_venta) as imp_venta, SUM(imp_venta_CUC) as imp_venta_CUC
    into #detalles     
    from TransDetalles 
    where cod_almacen = @cod_almacen
      and documento = @documento
      and numero = @numero
      and fecha = @fecha
      and cod_transaccion = @cod_transaccion;  
    
  update e
    set e.cantidad = d.cantidad,
        e.imp_costo = d.imp_costo,
        e.imp_venta = d.imp_venta,
        e.imp_venta_CUC = d.imp_venta_CUC
    from TransEncabezados e, #detalles d
    where e.cod_almacen = @cod_almacen
      and e.documento = @documento
      and e.numero = @numero
      and e.fecha = @fecha
      and e.cod_transaccion = @cod_transaccion;  

END
' 
END
GO
/****** Object:  StoredProcedure [dbo].[Utilitarios_Reconstruir_Encabezados]    Script Date: 12/04/2013 14:12:46 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Utilitarios_Reconstruir_Encabezados]') AND type in (N'P', N'PC'))
BEGIN
EXEC dbo.sp_executesql @statement = N'-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Utilitarios_Reconstruir_Encabezados]
AS
BEGIN

  select cod_almacen, documento, numero, fecha, cod_transaccion,
         SUM(cantidad) as cantidad, SUM(imp_costo) as imp_costo, SUM(imp_venta) as imp_venta, SUM(imp_venta_CUC) as imp_venta_CUC
    into #detalles     
    from TransDetalles 
    group by cod_almacen, documento, numero, fecha, cod_transaccion
  
  update e
    set e.cantidad = d.cantidad,
        e.imp_costo = d.imp_costo,
        e.imp_venta = d.imp_venta,
        e.imp_venta_CUC = d.imp_venta_CUC
    from TransEncabezados e, #detalles d
    where e.cod_almacen = d.cod_almacen
      and e.documento = d.documento
      and e.numero = d.numero
      and e.fecha = d.fecha
      and e.cod_transaccion = d.cod_transaccion;  
      
END
' 
END
GO
